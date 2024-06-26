package globals

import (
	"context"
	"github.com/PretendoNetwork/nex-go/v2/types"

	pb "github.com/PretendoNetwork/grpc-go/account"
	"github.com/PretendoNetwork/nex-go/v2"
	"google.golang.org/grpc/metadata"
)

func PasswordFromPID(pid *types.PID) (string, uint32) {
	ctx := metadata.NewOutgoingContext(context.Background(), GRPCAccountCommonMetadata)

	response, err := GRPCAccountClient.GetNEXPassword(ctx, &pb.GetNEXPasswordRequest{Pid: pid.LegacyValue()})
	if err != nil {
		Logger.Error(err.Error())
		return "", nex.ResultCodes.RendezVous.InvalidUsername
	}

	return response.Password, 0
}
