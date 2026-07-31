package database

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/splatoon/globals"
	ranking_splatoon_types "github.com/PretendoNetwork/splatoon/nex/ranking/types"
)

func StoreUserMatchResult(userPID types.PID, data ranking_splatoon_types.CompetitionRankingUploadScoreParam) error {
	var err error
	_, err = globals.Postgres.Exec(`
		INSERT INTO ranking_splatoon.results(
			uploader_pid,
			splatfest_id,
			score,
			team_id,
			team_score
		) VALUES ($1, $2, $3, $4, $5)
	`, int64(userPID), int64(data.SplatfestId), int64(data.Score), int16(data.TeamId), int64(data.TeamScore))

	if err != nil {
		return err
	}
	return nil
}
