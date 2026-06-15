package globals

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/splatoon"
)

func UploadCompetitionRankingScore(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		return nil, nex.NewError(nex.ResultCodes.Core.Unknown, err.Error())
	}

	request := packet.RMCMessage()
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	// Read CompetitionRankingUploadScoreParam data
	festivalID, streamErr := parametersStream.ReadUInt32LE()
	if streamErr != nil {
		return nil, nex.NewError(nex.ResultCodes.Core.Unknown, "Failed to read festival_id")
	}

	score, streamErr := parametersStream.ReadUInt32LE()
	if streamErr != nil {
		return nil, nex.NewError(nex.ResultCodes.Core.Unknown, "Failed to read score")
	}

	teamID, streamErr := parametersStream.ReadUInt8()
	if streamErr != nil {
		return nil, nex.NewError(nex.ResultCodes.Core.Unknown, "Failed to read team_id")
	}

	teamScore, streamErr := parametersStream.ReadUInt32LE()
	if streamErr != nil {
		return nil, nex.NewError(nex.ResultCodes.Core.Unknown, "Failed to read team_score")
	}

	isFirstUpload, streamErr := parametersStream.ReadBool()
	if streamErr != nil {
		return nil, nex.NewError(nex.ResultCodes.Core.Unknown, "Failed to read is_first_upload")
	}

	pid := packet.Sender().PID()

	_, dbErr := Postgres.Exec(
		"INSERT INTO competition_ranking_scores (pid, festival_id, score, team_id, team_score, is_first_upload) VALUES ($1, $2, $3, $4, $5, $6)",
		pid, festivalID, score, teamID, teamScore, isFirstUpload,
	)
	if dbErr != nil {
		Logger.Error(dbErr.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.Unknown, "Database error")
	}

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = ranking.ProtocolID
	rmcResponse.MethodID = ranking.MethodUploadCompetitionRankingScore
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
