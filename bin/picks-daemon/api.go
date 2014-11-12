package main

import (
	"net/rpc"
)

type Nil struct{}

func RegisterAPI(api interface{}) {
	rpc.Register(api)
}
