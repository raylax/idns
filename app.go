package main

import (
	"./util"
	"fmt"
	"net"
	"os"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:53")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	for {
		buf := make([]byte, 128)
		n, rAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		go process(rAddr, buf[:n])
	}
}

func process(rAddr *net.UDPAddr, data []byte) {
	util.ReadPacket(data)
}
