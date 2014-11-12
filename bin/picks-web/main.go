package main

import (
	"flag"
)

var (
	ListenAddr = flag.String("listen", ":10000", "Listen address")
)

func main() {
	flag.Parse()
}
