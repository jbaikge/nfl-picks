package picks

func (s *Store) Pick(userId int, p *Pick) (err error) {
	query := `INSERT OR REPLACE INTO picks
		(user_id, game_id, pick_value)
		VALUES
		(?,       ?,       ?         )
	`
	_, err = s.db.Exec(query, userId, p.GameId, p.Value)
	return
}
