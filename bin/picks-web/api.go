package main

import (
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
)

type Nil struct{}

var rpcServer = rpc.NewServer()

func init() {
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	http.Handle("/rpc", rpcServer)
}

func RegisterAPI(api interface{}) {
	if err := rpcServer.RegisterTCPService(api, ""); err != nil {
		log.Fatalf("rpcServer.RegisterService: %s", err)
	}
}
