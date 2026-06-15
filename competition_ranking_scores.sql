CREATE TABLE IF NOT EXISTS competition_ranking_scores (
    id bigserial PRIMARY KEY,
    pid numeric(10) NOT NULL,
    festival_id integer NOT NULL,
    score integer NOT NULL,
    team_id smallint NOT NULL,
    team_score integer NOT NULL,
    is_first_upload boolean NOT NULL
);
