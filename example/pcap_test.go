package example

import (
	"log"
	"testing"

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

func TestPcap(t *testing.T) {
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
}
