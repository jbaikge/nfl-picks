package picks

import (
	"database/sql"
	"time"
)

func (s *Store) FirstGameTime(w Week) (t time.Time, err error) {
	query := `SELECT
			game_start
		FROM games
		WHERE
			game_week     = $1
			AND game_year = $2
		ORDER BY game_start ASC
		LIMIT 1
	`
	err = s.db.QueryRow(query, w.Week, w.Year).Scan(&t)
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

func (s *Store) Scores(w Week) (games []*Game, err error) {
	query := `SELECT
			game_id,
			game_start,
			game_score_away,
			game_score_home,
			game_quarter,
			game_timeleft,
			game_posession
		FROM games
			LEFT JOIN stadiums USING(stadium_id)
		WHERE
			game_week     = $1
			AND game_year = $2
		ORDER BY game_start ASC, team_id_home ASC
	`
	rows, err := s.db.Query(query, w.Week, w.Year)
	if err != nil {
		return
	}
	defer rows.Close()

	games = make([]*Game, 0, 16)
	for rows.Next() {
		var gameId, quarter string
		var timeleft, posession sql.NullString
		g := new(Game)
		err = rows.Scan(
			&gameId,
			&g.Start,
			&g.AwayScore,
			&g.HomeScore,
			&quarter,
			&timeleft,
			&posession,
		)
		g.Id = GameIdType(gameId)
		g.Quarter = Quarter(quarter)
		// Convert the Postgres representation of an "interval" into a Go
		// time.Duration
		if timeleft.Valid {
			t, _ := time.Parse("15:04:05", timeleft.String)
			g.TimeLeft = time.Duration(t.AddDate(1970, 0, 0).UnixNano())
		}
		g.Posession = posession.String
		games = append(games, g)
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
	_, err = s.db.Exec(query, g.Id.String(), g.AwayScore, g.HomeScore, g.Posession, g.Quarter.Value(), g.TimeLeft.String())
	return
}
