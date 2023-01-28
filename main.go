package main

import (
	"fmt"
	"log"
	"time"
	"net"
	"golang.org/x/net/icmp"
)

func main() {
	// destUrl := "142.250.191.68" // www.google.com ip address

	conn, err := icmp.ListenPacket("udp4", "aaaa:55555")
	if err != nil {
		log.Fatal(err)
	}
	conn.SetDeadline(time.Now().Add(time.Minute))
	defer conn.Close()

	em := icmp.Echo {
		ID: 123,
		Seq: 1,
		Data: []byte("hello"),
	}

	eb, err := em.Marshal(0)
	if err != nil {
		log.Fatal(err)
	}

	destAddr, err:= net.ResolveUDPAddr("upd4", "https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.WriteTo(eb, destAddr)
	if err != nil {
		log.Fatal(err)
	}

	outByte := make([]byte, 1024)
	_, readAddr, err := conn.ReadFrom(outByte)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read from %v", readAddr.String())
}
