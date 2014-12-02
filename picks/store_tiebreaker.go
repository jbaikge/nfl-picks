package picks

import (
	"database/sql"
)

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

func (s *Store) AllTieBreakers(w Week) (tiebreakers map[string]float64, err error) {
	query := `SELECT
			user_name,
			tie_value
		FROM
			tie_breakers
			LEFT JOIN users USING(user_id)
		WHERE
			tie_week     = $1
			AND tie_year = $2
	`
	rows, err := s.db.Query(query, w.Week, w.Year)
	if err != nil {
		return
	}

	tiebreakers = make(map[string]float64)
	var username string
	var value float64
	for rows.Next() {
		if err = rows.Scan(&username, &value); err != nil {
			return
		}
		tiebreakers[username] = value
	}
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
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
