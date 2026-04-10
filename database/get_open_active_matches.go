package database

import (
	"database/sql"
	"time"

	pb "github.com/PretendoNetwork/grpc/go/nex/matchmaking/v1"
	"github.com/PretendoNetwork/splatoon/globals"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetOpenActiveMatches() (*[]pb.ActiveMatch, error) {

	rows, err := globals.Postgres.Query(`
	SELECT g.id, g.started_time, g.participants,g.owner_pid,s.game_mode,g.flags
	FROM matchmaking.gatherings g
	JOIN matchmaking.matchmake_sessions s ON (g.id = s.id)
	WHERE array_length(g.participants, 1) > 4 AND s.open_participation = true
	LIMIT 25`)

	if err != nil {
		if err == sql.ErrNoRows {
			globals.Logger.Error("Error no matches found")
			return nil, status.Errorf(codes.Internal, "internal server error")
		} else {
			globals.Logger.Error("Error unknown fetching matches")
			return nil, status.Errorf(codes.Internal, "internal server error")
		}
	}

	defer rows.Close()

	var matches []pb.ActiveMatch

	for rows.Next() {
		var id uint64
		var startTime time.Time
		var participants Participants
		var ownerPID uint32
		var gameMode uint64
		var flags uint64

		err := rows.Scan(&id, &startTime, &participants, &ownerPID, &gameMode, &flags)
		if err != nil {
			globals.Logger.Error("Error parsing row contents")
			globals.Logger.Error(err.Error())
			return nil, status.Errorf(codes.Internal, "internal server error")
		}

		matches = append(matches, pb.ActiveMatch{
			Id:           uint32(id),
			StartTime:    uint64(startTime.Unix()),
			Participants: participants,
			OwnerPid:     ownerPID,
			HostPid:      ownerPID,
			GameMode:     gameMode,
			Flags:        flags,
		})
	}
	return &matches, nil
}
