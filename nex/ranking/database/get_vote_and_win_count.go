package database

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/splatoon/globals"
)

// might be better to just return a struct here cause its not immediately obbvious from the signature which is which
func GetVoteAndWinCount(splatfestID uint32) (types.List[types.UInt32], types.List[types.UInt32], error) {
	row := globals.Postgres.QueryRow(`SELECT
		(SELECT
			COUNT(user_id)
		 FROM ranking_splatoon.user_scores WHERE splatfest_id=$1 AND team_id=0) as alpha_team_votes,
		(SELECT
			COUNT(user_id)
		 FROM ranking_splatoon.user_scores WHERE splatfest_id=$1 AND team_id=1) as bravo_team_votes,
		(SELECT
			SUM(team_score)
			FROM ranking_splatoon.results
			WHERE splatfest_id=$1
				AND team_id=1
		) as alpha_team_wins,
		(SELECT
			COUNT(team_score)
		 	FROM ranking_splatoon.results
			WHERE splatfest_id=$1
				AND team_id=1
		) as bravo_team_wins
	`, uint64(splatfestID))

	if row.Err() != nil {
		return types.NewList[types.UInt32](), types.NewList[types.UInt32](), row.Err()
	}
	var alphaTeamVotes uint32
	var bravoTeamVotes uint32
	var alphaTeamWins uint32
	var bravoTeamWins uint32

	err := row.Scan(&alphaTeamVotes, &bravoTeamVotes, &alphaTeamWins, &bravoTeamWins)
	if err != nil {
		return types.NewList[types.UInt32](), types.NewList[types.UInt32](), err
	}

	return types.List[types.UInt32]{types.NewUInt32(alphaTeamVotes), types.NewUInt32(bravoTeamVotes)},
		types.List[types.UInt32]{types.NewUInt32(alphaTeamWins), types.NewUInt32(bravoTeamWins)}, nil
}
