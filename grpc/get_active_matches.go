package grpc

import (
	"context"
	"fmt"

	pb "github.com/PretendoNetwork/grpc/go/nex/matchmaking/v1"
	db "github.com/PretendoNetwork/splatoon/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*gRPCNexServer) GetActiveMatches(ctx context.Context, in *pb.GetActiveMatchesRequest) (*pb.GetActiveMatchesResponse, error) {
	var result, err = db.GetOpenActiveMatches()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}
	var matches []*pb.ActiveMatch
	for i := range *result {
		element := &(*result)[i]
		fmt.Printf("ID: %v; Start Time: %v; PIDs: %v\n", element.Id, element.StartTime, element.Participants)
		matches = append(matches, element)
	}
	return &pb.GetActiveMatchesResponse{
		Matches: matches,
	}, nil
}
