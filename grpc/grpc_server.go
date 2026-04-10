package grpc

import (
	"fmt"
	"net"

	pb "github.com/PretendoNetwork/grpc/go/nex/matchmaking/v1"
	"github.com/PretendoNetwork/splatoon/globals"
	"google.golang.org/grpc"
)

type gRPCNexServer struct {
	pb.UnimplementedNEXMatchmakingServiceV1Server
}

func StartGRPCServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8123))
	if err != nil {
		globals.Logger.Errorf("failed to listen: %v", err)
	}

	server := grpc.NewServer(
	//grpc.UnaryInterceptor(apiKeyInterceptor),
	)

	pb.RegisterNEXMatchmakingServiceV1Server(server, &gRPCNexServer{})

	globals.Logger.Infof("GRPC server listening at %v", listener.Addr())

	if err := server.Serve(listener); err != nil {
		globals.Logger.Errorf("failed to serve: %v", err)
	}
}
