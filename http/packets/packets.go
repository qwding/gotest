package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
	"github.com/google/gopacket/layers"
)

var (
	device string = "en0"
	snapshot_len int32 = 1024
	promiscuous bool = false
	err error
	timeout time.Duration = 30 * time.Second
	handle       *pcap.Handle
)

func main() {
	fmt.Println("begin .... ")
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here

		if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
			//log.Println("Unusable packet")
			continue
		}

		tcp := packet.TransportLayer().(*layers.TCP)


		tcp.

		fmt.Printf("tcp %v\n:", string(tcp.LayerContents()))

		//fmt.Println(packet)
		//fmt.Println(string(packet.Data()))
	}
}
