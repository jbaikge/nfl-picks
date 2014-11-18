package picks

import (
	"database/sql"
	"fmt"
)

func (s *Store) UpdateUser(id int64, username string, pin int) (err error) {
	query := `UPDATE users
		SET
			user_name = $2,
			user_pin  = $3
		WHERE
			user_id   = $1
	`
	_, err = s.db.Exec(query, id, username, pin)
	return
}

func (s *Store) UserLastSeen(userId int64) (err error) {
	_, err = s.db.Exec(`UPDATE users SET user_last_seen = NOW() WHERE user_id = $1`, userId)
	return
}

func (s *Store) Usernames() (usernames []string, err error) {
	query := `SELECT user_name FROM users ORDER BY user_name ASC`
	rows, err := s.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	var username string
	usernames = make([]string, 0, 16)
	for rows.Next() {
		if err = rows.Scan(&username); err != nil {
			return
		}
		usernames = append(usernames, username)
	}
	return
}

func (s *Store) UserValidate(username string, pin int) (userId int64, isAdmin bool, err error) {
	query := `SELECT user_id, user_admin
		FROM users
		WHERE
			user_name = $1
			AND user_pin = $2
		LIMIT 1
	`
	err = s.db.QueryRow(query, username, fmt.Sprintf("%04d", pin)).Scan(&userId, &isAdmin)
	if err == sql.ErrNoRows {
		err = fmt.Errorf("Invalid credentials")
	}
	return
}

func (s *Store) UserValidatePIN(id int64, pin int) (match bool, err error) {
	query := `SELECT TRUE FROM users WHERE user_id = $1 AND user_pin = $2 LIMIT 1`
	err = s.db.QueryRow(query, id, fmt.Sprintf("%04d", pin)).Scan(&match)
	if err == sql.ErrNoRows {
		match = false
		err = nil
	}
	return
}
