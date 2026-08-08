package database

import (
	"github.com/PretendoNetwork/splatoon/globals"
	ranking_splatoon_types "github.com/PretendoNetwork/splatoon/nex/ranking/types"
)

func GetScoreDataEntries(splatfestID uint32) ([]ranking_splatoon_types.CompetitionRankingScoreData, error) {
	list := []ranking_splatoon_types.CompetitionRankingScoreData{}

	stream, err := globals.Postgres.Query(`SELECT
		uploader_pid,
		score,
		last_updated,
		app_data
	 FROM ranking_splatoon.user_scores WHERE splatfest_id=$1`, uint64(splatfestID))

	if err != nil {
		return nil, err
	}

	for stream.Next() {
		competitionRankingScoreData := ranking_splatoon_types.NewCompetitionRankingScoreData()
		err = stream.Scan(&competitionRankingScoreData.PID, &competitionRankingScoreData.Score, &competitionRankingScoreData.Modified, &competitionRankingScoreData.AppData)
		if err != nil {
			stream.Close()
			return nil, err
		}
		list = append(list, competitionRankingScoreData)
	}

	err = stream.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}
