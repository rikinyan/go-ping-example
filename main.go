package main

import (
	"fmt"
	"log"
	"time"
	"net"
	"golang.org/x/net/icmp"
)

func main() {
	conn, err := icmp.ListenPacket("udp4", "59.84.10.80:55555")
	if err != nil {
		log.Fatal(err)
	}
	conn.SetDeadline(time.Now().Add(time.Minute))
	defer conn.Close()
	fmt.Println("connection success")

	em := icmp.Echo {
		ID: 123,
		Seq: 1,
		Data: []byte(""),
	}

	eb, err := em.Marshal(0)
	if err != nil {
		log.Fatal(err)
	}

	destAddr, err:= net.ResolveUDPAddr("udp4", "142.251.42.142:80")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("resolve domain to ip addr")

	_, err = conn.WriteTo(eb, destAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("end to send")

	outByte := make([]byte, 1024)
	_, readAddr, err := conn.ReadFrom(outByte)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read from %v", readAddr.String())
}
