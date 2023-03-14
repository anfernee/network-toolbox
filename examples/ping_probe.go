package main

import (
	"log"

	"github.com/anfernee/network-toolbox/pkg/ip"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/icmp"
)

func main() {
	h := func(msg *icmp.Message, from string) {
		log.Printf("Received icmp from %v", from)
		log.Println(spew.Sdump(msg))
	}

	if err := ip.ListenICMP(ip.ICMPHandleFunc(h)); err != nil {
		log.Fatalf("listen icmp error: %v", err)
	}
}
