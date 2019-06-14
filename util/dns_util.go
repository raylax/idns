package util

import "net"

func QueryUpstream(upstream string, data []byte, bufSize int) ([]byte, error) {
	conn, err := net.Dial("udp", upstream+":53")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	_, err = conn.Write(data)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, bufSize)
	n, err := conn.Read(buf)
	return buf[:n], nil
}
