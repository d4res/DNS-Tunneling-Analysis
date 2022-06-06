package pcap

import (
	"fmt"
	"testing"
)

func TestPcap(t *testing.T) {
	source, handle, err := NewSource()
	defer handle.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("test1")
	for packet := range source.Packets() {
		fmt.Println(packet)
	}

	fmt.Println("test")
}
