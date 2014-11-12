package main

import (
	"github.com/gorilla/rpc"
	"log"
)

type Nil struct{}

var rpcServer = rpc.NewServer()

func RegisterAPI(api interface{}) {
	if err := rpcServer.RegisterService(receiver, ""); err != nil {
		log.Fatalf("rpcServer.RegisterService: %s", err)
	}
}
