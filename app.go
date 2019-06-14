package main

import (
	"./util"
	"log"
	"net"
)

const (
	// 1kb
	ClientReadSize = 1024
	// 1kb
	UpstreamReadSize = 1024
	// process pool size
	ProcessPoolSize = 1000
)

var processPoolCh = make(chan bool, ProcessPoolSize)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:53")
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.SetReadBuffer(ClientReadSize)
	for {
		processPoolCh <- true
		read(conn)
	}
}

func read(conn *net.UDPConn) {
	buf := make([]byte, ClientReadSize)
	for {
		n, rAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Failed to read data from client", err)
			continue
		}
		err = process(conn, rAddr, buf[:n])
		if err != nil {
			log.Println("Failed to process request", err)
			continue
		}
	}
	<-processPoolCh
}

func process(conn *net.UDPConn, rAddr *net.UDPAddr, data []byte) error {
	result, err := util.QueryUpstream("114.114.114.114", data, UpstreamReadSize)
	_ = util.ReadPacket(result)
	//println(packet.String())
	if err != nil {
		log.Println("Failed to query from upstream", err)
		return err
	}
	_, err = conn.WriteToUDP(result, rAddr)
	return err
}
