package picks

func (s *Store) AddTieBreaker(userId int64, t TieBreaker) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return
	}

	del := `DELETE FROM tie_breakers WHERE user_id = $1 AND tie_week = $2 AND tie_year = $3`
	if _, err = s.db.Exec(del, userId, t.Week.Week, t.Week.Year); err != nil {
		return
	}

	insert := `INSERT INTO tie_breakers
		(user_id, tie_week, tie_year, tie_value, tie_added)
		VALUES
		($1,      $2,       $3,       $4,        NOW()    )
	`
	if _, err = s.db.Exec(insert, userId, t.Week.Week, t.Week.Year, t.Value); err != nil {
		return
	}

	err = tx.Commit()
	return
}

func (s *Store) UserTieBreaker(userId int64, w Week) (t TieBreaker, err error) {
	query := `SELECT
			tie_week, tie_year, tie_value
		FROM tie_breakers
		WHERE
			user_id      = $1
			AND tie_week = $2
			AND tie_year = $3
	`
	row := s.db.QueryRow(query, userId, w.Week, w.Year)
	err = row.Scan(&t.Week.Week, &t.Week.Year, &t.Value)
	return
}
