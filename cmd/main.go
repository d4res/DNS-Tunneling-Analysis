package main

import (
	"DNSpcap/decode"
	"DNSpcap/model"
	"DNSpcap/proto"
	"context"
	"time"

	"github.com/google/gopacket/layers"

	pcapWrap "DNSpcap/pcap"
)

func main() {
	client, err := model.Conn()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())
	requestCol := client.Database("dns_pcap").Collection("request")
	responseCol := client.Database("dns_pcap").Collection("response")
	infoCol := client.Database("dns_pcap").Collection("info")

	// devices, err := pcap.FindAllDevs()
	// if err != nil {
	// 	log.Panicln(err)
	// }

	// for _, device := range devices {
	// 	if device.Name == iface {
	// 		devFound = true
	// 	}
	// }
	// if !devFound {
	// 	log.Panicf("Device named '%s' does not exist\n", iface)
	// }

	// handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// defer handle.Close()

	// if err := handle.SetBPFFilter(filter); err != nil {
	// 	log.Panicln(err)
	// }

	source, handle, err := pcapWrap.NewSource()
	defer handle.Close()
	if err != nil {
		panic(err)
	}
	//source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		dnsLayer := packet.Layer(layers.LayerTypeDNS)
		dns := dnsLayer.(*layers.DNS)

		// request 0; response 1
		if dns.QR {
			// this is response
			for _, v := range dns.Answers {
				var data model.Response
				data.Class = v.Class.String()
				data.Type = v.Type.String()
				data.Time = time.Now()
				var payload string
				var isEval bool

				// now considering TXT, MX, CNAME
				switch v.Type {
				case layers.DNSTypeTXT:
					payload = string(v.TXT)
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

				if isEval {
					msg, err := decode.ParseMsg(decode.TransDomain(data.Payload, "sub.dares.top"))

					if msg.Payload != nil && err == nil {
						infoCol.InsertOne(context.TODO(), msg)
					}
				}

				responseCol.InsertOne(context.TODO(), data)
			}
		} else {
			// this is request
			for _, v := range dns.Questions {

				tag := proto.IsEval(string(v.Name))
				data := model.Request{Class: v.Class.String(), Type: v.Type.String(), Payload: string(v.Name), Tag: tag, Time: time.Now()}

				if tag {
					msg, err := decode.ParseMsg(decode.TransDomain(data.Payload, "sub.dares.top"))

					if msg.Payload != nil && err == nil {
						infoCol.InsertOne(context.TODO(), msg)
					}
				}

				requestCol.InsertOne(context.TODO(), data)
			}
		}
	}
}
