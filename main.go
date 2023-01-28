package main

import (
	"fmt"
	"log"
	"time"
	"net"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func main() {
	conn, err := icmp.ListenPacket("udp4", "")
	if err != nil {
		log.Fatal(err)
	}
	conn.SetDeadline(time.Now().Add(time.Minute))
	defer conn.Close()
	fmt.Println("connection success")

	echo := icmp.Echo {
		ID: 123,
		Seq: 1,
		Data: []byte(""),
	}

	message := icmp.Message {
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &echo,
	}

	destAddr, err:= net.ResolveUDPAddr("udp4", "142.251.42.142:80")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("resolve domain to ip addr")

	mb, _ := message.Marshal(nil)
	_, err = conn.WriteTo(mb, destAddr)
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
