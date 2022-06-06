package pcap

import "github.com/google/gopacket/pcap"

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
