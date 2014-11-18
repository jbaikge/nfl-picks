package picks

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Store struct {
	dsn string
	db  *sql.DB
}

func NewStore(dataSourceName string) (s *Store, err error) {
	s = &Store{
		dsn: dataSourceName,
	}
	s.db, err = sql.Open("postgres", dataSourceName)
	return
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) CurrentWeek() (year, week int, season string, err error) {
	query := `SELECT year, week, season FROM config`
	err = s.db.QueryRow(query).Scan(&year, &week, &season)
	return
}

func (s *Store) NewGame(g *Game) (err error) {
	query := `INSERT INTO games
		(game_id, nfl_event_id, stadium_id, team_id_home, team_id_away, game_score_home, game_score_away, game_start, game_quarter, game_week, game_year, game_season)
		VALUES
		($1,      $2,           $3,         $4,           $5,           $6,              $7,              $8,         $9,           $10,       $11,       $12        )
	`
	_, err = s.db.Exec(query, g.Id.String(), g.EventId, g.Home, g.Home, g.Away, g.HomeScore, g.AwayScore, g.Start, string(g.Quarter), g.Week, g.Year, g.Season)
	return
}

func (s *Store) Pick(userId int, p *Pick) (err error) {
	query := `INSERT OR REPLACE INTO picks
		(user_id, game_id, pick_value)
		VALUES
		(?,       ?,       ?         )
	`
	_, err = s.db.Exec(query, userId, p.GameId, p.Value)
	return
}

func (s *Store) UpdateCurrentWeek(year, week int, season string) (err error) {
	query := `UPDATE config SET
		year   = $1,
		week   = $2,
		season = $3
	`
	_, err = s.db.Exec(query, year, week, season)
	return
}

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
	_, err = s.db.Exec(query, g.Id.String(), g.AwayScore, g.HomeScore, g.Posession, g.Quarter.Value(), g.TimeLeft)
	return
}

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
	query := `SELECT user_name FROM users`
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
