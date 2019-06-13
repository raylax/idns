package main

import (
	"./util"
	"encoding/json"
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
		process(rAddr, buf[:n])
	}
}

func process(rAddr *net.UDPAddr, data []byte) {
	packet := util.ReadPacket(data)
	json, err := json.Marshal(packet)
	if err != nil {
		panic(err)
	}
	println(string(json))
}
