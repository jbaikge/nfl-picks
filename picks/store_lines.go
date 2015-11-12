package picks

import (
	"fmt"
	"github.com/lib/pq"
)

// Couple human-friendly extras in PickLine over regular Line
func (s *Store) CurrentPickLines() (w Week, lines []*PickLine, err error) {
	if w, err = s.CurrentWeek(); err != nil {
		return
	}
	lines, err = s.PickLines(w)
	return
}

func (s *Store) PickLines(w Week) (lines []*PickLine, err error) {
	standingsQueryF := `SELECT
			team_id,
			COALESCE(COUNT(NULLIF(win, FALSE)), 0) AS wins,
			COALESCE(COUNT(NULLIF(win, TRUE)), 0) AS losses
		FROM
			(
				SELECT
					team_id_home AS team_id,
					game_year,
					game_week,
					game_score_home > game_score_away AS win
				FROM
					games
				UNION ALL
				SELECT
					team_id_away,
					game_year,
					game_week,
					game_score_away > game_score_home
				FROM games
			) AS t
		WHERE
			game_week     < %d
			AND game_year = %d
		GROUP BY team_id
	`
	standingsQuery := fmt.Sprintf(standingsQueryF, w.Week, w.Year)
	query := `SELECT
			game_start,
			game_id,
			game_spread,
			game_over_under,
			game_line_updated,
			homeTeam.team_id,
			homeTeam.team_city,
			homeTeam.team_name,
			homeTeam.team_league,
			homeTeam.team_division,
			homeStandings.wins,
			homeStandings.losses,
			awayTeam.team_id,
			awayTeam.team_city,
			awayTeam.team_name,
			awayTeam.team_league,
			awayTeam.team_division,
			awayStandings.wins,
			awayStandings.losses,
			stadiums.stadium_id,
			stadiums.stadium_name,
			stadiums.stadium_city,
			stadiums.stadium_state,
			stadiums.stadium_turf,
			stadiums.stadium_roof
		FROM games
			LEFT JOIN stadiums USING(stadium_id)
			LEFT JOIN teams AS homeTeam ON(games.team_id_home = homeTeam.team_id)
			LEFT JOIN teams AS awayTeam ON(games.team_id_away = awayTeam.team_id)
			LEFT JOIN (` + standingsQuery + `) AS homeStandings ON(games.team_id_home = homeStandings.team_id)
			LEFT JOIN (` + standingsQuery + `) AS awayStandings ON(games.team_id_away = awayStandings.team_id)
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

	lines = make([]*PickLine, 0, 16)
	for rows.Next() {
		var game_id string
		var updated pq.NullTime
		l := new(PickLine)
		err = rows.Scan(
			&l.Start,
			&game_id,
			&l.Line.Spread,
			&l.Line.OverUnder,
			&updated,
			&l.Home.Id,
			&l.Home.City,
			&l.Home.Name,
			&l.Home.League,
			&l.Home.Division,
			&l.Home.Wins,
			&l.Home.Losses,
			&l.Away.Id,
			&l.Away.City,
			&l.Away.Name,
			&l.Away.League,
			&l.Away.Division,
			&l.Away.Wins,
			&l.Away.Losses,
			&l.Stadium.Id,
			&l.Stadium.Name,
			&l.Stadium.City,
			&l.Stadium.State,
			&l.Stadium.Turf,
			&l.Stadium.Roof,
		)
		l.Line.GameId = GameIdType(game_id)
		if updated.Valid {
			l.Line.Updated = updated.Time
		}
		lines = append(lines, l)
	}
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
