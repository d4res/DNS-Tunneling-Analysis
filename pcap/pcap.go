package pcap

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func NewSource() (*gopacket.PacketSource, *pcap.Handle, error) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return nil, nil, err
	}

	for _, device := range devices {
		if device.Name == iface {
			devFound = true
		}
	}

	if !devFound {
		return nil, nil, fmt.Errorf("device named '%s' does not exist", iface)
	}

	handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
	if err != nil {
		return nil, nil, err
	}
	//defer handle.Close()

	if err := handle.SetBPFFilter(filter); err != nil {
		return nil, nil, err
	}

	return gopacket.NewPacketSource(handle, handle.LinkType()), handle, nil
}
