package picks

import (
	"fmt"
)

func (s *Store) UpdateLine(line *Line) (err error) {
	query := `UPDATE games
		SET
			game_spread       = $2,
			game_over_under   = $3,
			game_line_updated = $4
		WHERE
			game_id           = $1
	`
	result, err := s.db.Exec(query, line.GameId.String(), line.Spread, line.OverUnder, line.Updated)
	if err != nil {
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		return
	}
	if n != 1 {
		return fmt.Errorf("Store.UpdateLine: Expected to update one row, updated %d", n)
	}
	return
}
