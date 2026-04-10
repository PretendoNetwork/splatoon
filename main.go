package main

import (
	"sync"

	"github.com/PretendoNetwork/splatoon/grpc"
	"github.com/PretendoNetwork/splatoon/nex"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	// TODO - Add gRPC server
	go nex.StartAuthenticationServer()
	go nex.StartSecureServer()
	go grpc.StartGRPCServer()
	wg.Wait()
}
