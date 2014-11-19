package picks

func (s *Store) AllPicks(c Current) (picks map[GameIdType]map[string]Pick, err error) {
	query := `SELECT
			picks.game_id,
			picks.pick_value,
			users.user_name
		FROM
			picks
			LEFT JOIN games USING(game_id)
			LEFT JOIN users USING(user_id)
		WHERE
			games.game_week     = $1
			AND games.game_year = $2
	`

	rows, err := s.db.Query(query, c.Week, c.Year)
	if err != nil {
		return
	}
	defer rows.Close()

	var gameId string
	var userName string
	picks = make(map[GameIdType]map[string]Pick)
	for rows.Next() {
		var p Pick
		if err = rows.Scan(&gameId, &p.Value, &userName); err != nil {
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

func (s *Store) Pick(userId int64, p *Pick) (err error) {
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

func (s *Store) UserPicks(userId int64, c Current) (picks []*Pick, err error) {
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
	rows, err := s.db.Query(query, userId, c.Week, c.Year)
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
