package picks

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
		(user_id, game_id, pick_value)
		VALUES
		($1,      $2,      $3        )`
	if _, err = tx.Exec(ins, userId, p.GameId.String(), p.Value); err != nil {
		return
	}
	err = tx.Commit()
	return
}
