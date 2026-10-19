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

func GetSingleCompetitionRankingScore(splatfestID uint32) (ranking_splatoon_types.CompetitionRankingScoreInfo, error) {
	var err error
	info := ranking_splatoon_types.NewCompetitionRankingScoreInfo()

	info.FestID = types.UInt32(splatfestID)

	teamWins, teamVotes, err := database.GetVoteAndWinCount(splatfestID)

	if err != nil {
		return info, err
	}

	info.TeamWins = types.List[types.UInt32](teamWins)
	info.TeamVotes = types.List[types.UInt32](teamVotes)

	var scoreDataEntries []ranking_splatoon_types.CompetitionRankingScoreData
	scoreDataEntries, err = database.GetScoreDataEntries(splatfestID)
	if err != nil {
		return info, err
	}
	info.ScoreData = scoreDataEntries

	return info, nil
}

func GetCompetitionRankingScore(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) {
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	parameters := packet.RMCMessage().Parameters
	parametersStream := nex.NewByteStreamIn(parameters, globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	params := ranking_splatoon_types.NewCompetitionRankingGetParam()

	err = params.ExtractFrom(parametersStream)
	if err != nil {
		common_globals.Logger.Error("Failed to extract param on call to GetCompetitionRankingScore.")
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	retVal := types.NewList[ranking_splatoon_types.CompetitionRankingScoreInfo]()

	festID := params.ResultRange.Offset
	info, err := GetSingleCompetitionRankingScore(uint32(festID))
	if err != nil {
		return nil, nex.NewError(nex.ResultCodes.Core.Exception, "error retrieving ranking scores")
	}
	retVal = append(retVal, info)

	retVal.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = ranking.ProtocolID
	rmcResponse.MethodID = ranking.MethodGetCompetitionRankingScore
	rmcResponse.CallID = callID

	return rmcResponse, nil

}
