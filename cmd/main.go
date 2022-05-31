package main

import (
	"DNSpcap/model"
	"DNSpcap/proto"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	iface   = "enp0s3"
	snaplen = int32(1600)
	promisc = false
	timeout = pcap.BlockForever
	// setting the BPF filter, look up: https://www.ibm.com/docs/en/qsip/7.4?topic=queries-berkeley-packet-filters.
	// here, we shall capture all udp packets through port 53
	filter   = "udp and port 53"
	devFound = false
)

func main() {
	client, err := model.Conn()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())
	requestCol := client.Database("dns_pcap").Collection("request")
	responseCol := client.Database("dns_pcap").Collection("response")

	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}

	for _, device := range devices {
		if device.Name == iface {
			devFound = true
		}
	}
	if !devFound {
		log.Panicf("Device named '%s' does not exist\n", iface)
	}

	handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
	if err != nil {
		log.Panicln(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(filter); err != nil {
		log.Panicln(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		// fmt.Println(packet.Dump())
		//fmt.Println(packet)

		// DNS is specified in RFC 1034 / RFC 1035
		// +---------------------+
		// |        Header       |
		// +---------------------+
		// |       Question      | the question for the name server
		// +---------------------+
		// |        Answer       | RRs answering the question
		// +---------------------+
		// |      Authority      | RRs pointing toward an authority
		// +---------------------+
		// |      Additional     | RRs holding additional information
		// +---------------------+
		//
		//  DNS Header
		//  0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
		//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		//  |                      ID                       |
		//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		//  |QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
		//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		//  |                    QDCOUNT                    |
		//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		//  |                    ANCOUNT                    |
		//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		//  |                    NSCOUNT                    |
		//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		//  |                    ARCOUNT                    |
		//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		dnsLayer := packet.Layer(layers.LayerTypeDNS)
		dns := dnsLayer.(*layers.DNS)

		// request 0; response 1
		if dns.QR {
			for _, v := range dns.Answers {
				///fmt.Println(v)

				var data model.Response
				data.Class = v.Class.String()
				data.Type = v.Type.String()
				data.Time = time.Now()
				var payload string
				var isEval bool

				// now considering TXT, MX, CNAME
				switch v.Type {
				case layers.DNSTypeA:
					payload = v.IP.String()
				case layers.DNSTypeCNAME:
					payload = string(v.CNAME)
				case layers.DNSTypeMX:
					payload = string(v.MX.Name)
				default:
					continue
				}

				isEval = proto.IsEval(payload)
				data.Tag = isEval
				data.Payload = payload

				responseCol.InsertOne(context.TODO(), data)
				//fmt.Printf("%s %s %s\n", v.)
			}
		} else {

			for _, v := range dns.Questions {
				fmt.Println("request", string(v.Name))
				tag := proto.IsEval(string(v.Name))
				data := model.Request{Class: v.Class.String(), Type: v.Type.String(), Payload: string(v.Name), Tag: tag, Time: time.Now()}
				requestCol.InsertOne(context.TODO(), data)
			}
		}
	}
}
