package database

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/splatoon/globals"
	ranking_splatoon_types "github.com/PretendoNetwork/splatoon/nex/ranking/types"
)

func StoreUserScore(userPID types.PID, data ranking_splatoon_types.CompetitionRankingUploadScoreParam) error {
	var err error
	time := types.NewDateTime(0).Now()
	_, err = globals.Postgres.Exec(`
		INSERT INTO ranking_splatoon.user_scores(
			uploader_pid,
			splatfest_id,
			score,
			team_id,
			last_updated,
			app_data
		) VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (uploader_pid, splatfest_id)
		DO UPDATE SET
			score = EXCLUDED.score,
			last_updated = EXCLUDED.last_updated
			team_id = EXCLUDED.team_id
			app_data = EXCLUDED.app_data
	`, int64(userPID), int64(data.SplatfestId), int64(data.Score), int16(data.TeamID), int64(time), ([]byte)(data.AppData))

	if err != nil {
		return err
	}
	return nil
}
