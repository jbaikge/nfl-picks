package picks

func (s *Store) AddPick(userId int64, p *Pick) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return
	}
	del := `DELETE FROM picks WHERE user_id = $1 AND game_id = $2`
	if _, err = tx.Exec(del, userId, p.GameId.String()); err != nil {
		return
	}
	ins := `INSERT INTO picks
		(user_id, game_id, pick_value, pick_added)
		VALUES
		($1,      $2,      $3,         NOW()     )`
	if _, err = tx.Exec(ins, userId, p.GameId.String(), p.Value); err != nil {
		return
	}
	err = tx.Commit()
	return
}

func (s *Store) AllPicks(w Week) (picks map[GameIdType]map[string]Pick, err error) {
	query := `SELECT
			picks.game_id,
			picks.pick_value,
			users.user_name,
			(
				picks.pick_value = games.team_id_away
				AND games.game_score_away - games.game_spread >= games.game_score_home
			) OR (
				picks.pick_value = games.team_id_away
				AND games.game_score_home + games.game_spread >= games.game_score_away
			) OR (
				picks.pick_value = 'UNDER'
				AND games.game_score_away + games.game_score_home <= games.game_over_under
			) OR (
				picks.pick_value = 'OVER'
				AND games.game_score_away + games.game_score_home >= games.game_over_under
			) OR (
				games.game_quarter = 'P'
			) AS correct
		FROM
			picks
			LEFT JOIN games USING(game_id)
			LEFT JOIN users USING(user_id)
		WHERE
			games.game_week     = $1
			AND games.game_year = $2
	`

	rows, err := s.db.Query(query, w.Week, w.Year)
	if err != nil {
		return
	}
	defer rows.Close()

	var gameId string
	var userName string
	picks = make(map[GameIdType]map[string]Pick)
	for rows.Next() {
		var p Pick
		if err = rows.Scan(&gameId, &p.Value, &userName, &p.Correct); err != nil {
			return
		}
		p.GameId = GameIdType(gameId)
		if picks[p.GameId] == nil {
			picks[p.GameId] = make(map[string]Pick)
		}
		picks[p.GameId][userName] = p
	}

	return
}

func (s *Store) UserPicks(userId int64, w Week) (picks []*Pick, err error) {
	query := `SELECT
			game_id,
			pick_value
		FROM
			picks
			LEFT JOIN games USING(game_id)
		WHERE
			picks.user_id       = $1
			AND games.game_week = $2
			AND games.game_year = $3
	`
	rows, err := s.db.Query(query, userId, w.Week, w.Year)
	if err != nil {
		return
	}
	defer rows.Close()

	var gameId string
	picks = make([]*Pick, 0, 16)
	for rows.Next() {
		p := new(Pick)
		if err = rows.Scan(&gameId, &p.Value); err != nil {
			return
		}
		p.GameId = GameIdType(gameId)
		picks = append(picks, p)
	}
	return
}

// SELECT game_id, game_spread AS s, game_over_under AS ou,
// team_id_away AS a,
// game_score_away AS as,
// game_score_away - game_spread AS ad,
// game_score_away - game_spread >= game_score_home AS awin,
// team_id_home AS h,
// game_score_home AS hs,
// game_score_home + game_spread AS hd,
// game_score_home + game_spread >= game_score_away AS hwin,
// game_score_away + game_score_home >= game_over_under AS over,
// game_score_away + game_score_home <= game_over_under AS under
// FROM games WHERE game_week = 1 AND game_year = 2014;
// func (s *Store) WinningPicks(w Week) (picks []GamePick, err error) {
// 	query := `SELECT
// 			game_id,
// 			team_id_away,
// 			game_score_away - game_spread >= game_score_home,
// 			team_id_home,
// 			game_score_home + game_spread >= game_score_away,
// 			game_score_away + game_score_home >= game_over_under,
// 			game_score_away + game_score_home <= game_over_under
// 		FROM
// 			games
// 		WHERE
// 			games.game_week     = $1
// 			AND games.game_year = $2
// 	`
// 	rows, err := s.db.Query(query, w.Week, w.Year)
// 	if err != nil {
// 		return
// 	}

// 	var gameId, awayId, homeId string
// 	var away, home, over, under bool
// 	picks = make([]GamePick, 0, 16)
// 	for rows.Next() {
// 		var p GamePick
// 		err = rows.Scan(&gameId, &p.AwayId, &p.Away, &p.HomeId, &p.Home, &p.Over, &p.Under)
// 		if err != nil {
// 			return
// 		}
// 		p.GameId = GameIdType(gameId)
// 		picks = append(picks, p)
// 	}
// 	return
// }
