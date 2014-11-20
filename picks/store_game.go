package picks

func (s *Store) NewGame(g *Game) (err error) {
	query := `INSERT INTO games
		(game_id, nfl_event_id, stadium_id, team_id_home, team_id_away, game_score_home, game_score_away, game_start, game_quarter, game_week, game_year, game_season)
		VALUES
		($1,      $2,           $3,         $4,           $5,           $6,              $7,              $8,         $9,           $10,       $11,       $12        )
	`
	_, err = s.db.Exec(query, g.Id.String(), g.EventId, g.Home, g.Home, g.Away, g.HomeScore, g.AwayScore, g.Start, string(g.Quarter), g.Week, g.Year, g.Season)
	return
}

func (s *Store) UpdateGame(g *Game) (err error) {
	query := `UPDATE games
		SET
			game_score_away    = $2,
			game_score_home    = $3,
			game_posession     = $4,
			game_quarter       = $5,
			game_timeleft      = $6,
			game_score_updated = NOW()
		WHERE
			game_id            = $1
	`
	_, err = s.db.Exec(query, g.Id.String(), g.AwayScore, g.HomeScore, g.Posession, g.Quarter.Value(), g.TimeLeft.String())
	return
}
