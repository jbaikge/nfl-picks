package picks

import (
	"database/sql"
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

func (s *Store) CurrentWeek() (c Current, err error) {
	query := `SELECT year, week, season FROM config`
	err = s.db.QueryRow(query).Scan(&c.Year, &c.Week, &c.Season)
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
