package main

import (
	"net"
	"log"
	"flag"
	"net/rpc"
	"fmt"
)

var (
	ListenAddr = flag.String("listen", ":10000", "Listen address")
)


func main() {
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", *ListenAddr)
	if err != nil {
		log.Fatalf("net.Listen: %s", err)
	}
	fmt.Printf("Accepting connections at %s\n", *ListenAddr)
	http.Serve(l, nil)
}
