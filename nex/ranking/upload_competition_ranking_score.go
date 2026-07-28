package nex_ranking_splaton

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/splatoon"
	"github.com/PretendoNetwork/splatoon/globals"
	"github.com/PretendoNetwork/splatoon/nex/ranking/database"
	ranking_splatoon_types "github.com/PretendoNetwork/splatoon/nex/ranking/types"
)

func UploadCompetitionRankingScore(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) {
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	parameters := packet.RMCMessage().Parameters
	parametersStream := nex.NewByteStreamIn(parameters, globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	params := ranking_splatoon_types.NewCompetitionRankingUploadScoreParam()

	err = params.ExtractFrom(parametersStream)
	if err != nil {
		common_globals.Logger.Error("Failed to extract param on call to UploadCompetitionRankingScore.")
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	database.StoreUserMatchResult(packet.Sender().PID(), params)
	database.StoreUserScore(packet.Sender().PID(), params)

	types.NewBool(true).WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = ranking.ProtocolID
	rmcResponse.MethodID = ranking.MethodGetCompetitionRankingScore
	rmcResponse.CallID = callID

	return rmcResponse, nil

}
