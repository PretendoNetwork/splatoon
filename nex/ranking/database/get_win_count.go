package database

import (
	"github.com/PretendoNetwork/splatoon/globals"
)

func GetWinCount(splatfestID uint32, team_id uint8) (uint32, error) {
	row := globals.Postgres.QueryRow(`SELECT
			COUNT(upload_id)
	 	FROM ranking_splatoon.results
		WHERE splatfest_id=$1
			AND (
			(has_won = false AND team_id NOT EQUAL $2) OR
			(has_won = true AND team_id=$2)
		)`, uint64(splatfestID), int16(team_id))

	if row.Err() != nil {
		return 0, row.Err()
	}
	var count uint32

	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
