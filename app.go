package main

import (
	"./util"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	ClientReadSize   = 1024
	UpstreamReadSize = 1024
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

	conn.SetReadBuffer(ClientReadSize)
	buf := make([]byte, ClientReadSize)

	for {
		start := time.Now()
		n, rAddr, err := conn.ReadFromUDP(buf)
		util.LogSlowStepTimes(fmt.Sprintf("Read from client size:%v", n), start)
		if err != nil {
			continue
		}
		go process(conn, rAddr, buf[:n])
	}
}

func process(conn *net.UDPConn, rAddr *net.UDPAddr, data []byte) {
	packet := util.ReadPacket(data)
	_, err := json.Marshal(packet)
	if err != nil {
		panic(err)
	}

	result := queryUpstream(data)
	packet = util.ReadPacket(data)

	start := time.Now()
	n, err := conn.WriteToUDP(result, rAddr)
	util.LogSlowStepTimes(fmt.Sprintf("Write to client size:%v", n), start)
	if err != nil {
		log.Println(err)
	}
}

func queryUpstream(data []byte) []byte {
	start := time.Now()
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	_, err = conn.Write(data)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, UpstreamReadSize)
	n, err := conn.Read(buf)
	util.LogSlowStepTimes(fmt.Sprintf("Query from upstream sent:%v recv:%v", len(data), n), start)
	return buf[:n]
}
