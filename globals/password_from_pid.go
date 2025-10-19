package globals

import (
	"context"
	"encoding/json"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	"os"
	"strconv"

	pb "github.com/PretendoNetwork/grpc/go/account"
	"github.com/PretendoNetwork/nex-go/v2"
	"google.golang.org/grpc/metadata"
)

func PasswordFromPID(pid types.PID) (string, uint32) {
	ctx := metadata.NewOutgoingContext(context.Background(), GRPCAccountCommonMetadata)

	response, err := GRPCAccountClient.GetNEXPassword(ctx, &pb.GetNEXPasswordRequest{Pid: uint32(pid)})
	if err != nil {
		Logger.Error(err.Error())
		return "", nex.ResultCodes.RendezVous.InvalidUsername
	}

	return response.Password, 0
}

// This is the same format as nex-viewer's settings.json
type jsonAccount struct {
	Platform string  `json:"platform"`
	Username string  `json:"username"`
	Pid      float64 `json:"pid"`
	Password string  `json:"password"`
}

type settingsJson struct {
	Accounts []jsonAccount `json:"accounts"`
}

// PasswordFromPIDLocal is an alternative NEX password validator that can be used offline
func PasswordFromPIDLocal(pid types.PID) (string, uint32) {
	file, err := os.ReadFile("settings.json")
	if err != nil {
		Logger.Error(err.Error())
		return "", nex.ResultCodes.RendezVous.InvalidUsername
	}

	var data *settingsJson
	err = json.Unmarshal(file, &data)
	if err != nil {
		Logger.Error(err.Error())
		return "", nex.ResultCodes.RendezVous.InvalidUsername
	}

	for _, account := range data.Accounts {
		if account.Username == strconv.FormatUint(uint64(pid), 10) {
			globals.Logger.Infof("Using local account details for %v", account.Username)
			return account.Password, 0
		}
	}

	return "", nex.ResultCodes.RendezVous.InvalidUsername
}
