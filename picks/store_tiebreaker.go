package picks

func (s *Store) AddTieBreaker(userId int64, t TieBreaker) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return
	}

	del := `DELETE FROM tie_breakers WHERE user_id = $1 AND game_id = $2`
	if _, err = s.db.Exec(del, userId, t.GameId.String()); err != nil {
		return
	}

	insert := `INSERT INTO tie_breakers
		(user_id, game_id, tie_value, tie_added)
		VALUES
		($1,      $2,      $3,        NOW()    )
	`
	if _, err = s.db.Exec(insert, userId, t.GameId.String(), t.Value); err != nil {
		return
	}

	err = tx.Commit()
	return
}
