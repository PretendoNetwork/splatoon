package database

import (
	"github.com/PretendoNetwork/splatoon/globals"
)

func Setup() {
	var err error

	_, err = globals.Postgres.Exec(`CREATE SCHEMA IF NOT EXISTS ranking_splatoon`)
	if err != nil {
		globals.Logger.Error(err.Error())
		return
	}

	_, err = globals.Postgres.Exec(`CREATE TABLE IF NOT EXISTS ranking_splatoon.results (
		upload_id bigserial primary key not null,
		uploader_pid int8 not null,
		splatfest_id int8 not null,
		team_id int2 not null,
		has_won int2 not null,
		app_data bytea not null,
	)`)
	_, err = globals.Postgres.Exec(`CREATE TABLE IF NOT EXISTS ranking_splatoon.user_scores (
		uploader_pid int8 not null,
		splatfest_id int8 not null,
		score int8 not null,
		team_id boolean not null,
		user_data bytea not null,
		last_updated int8 not null
		PRIMARY KEY(uploader_pid, splatfest_id)
	)`)
	if err != nil {
		globals.Logger.Error(err.Error())
		return
	}
}
