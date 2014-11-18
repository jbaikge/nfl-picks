package picks

import (
	"fmt"
)

// Couple human-friendly extras in PickLine over regular Line
func (s *Store) CurrentPickLines() (c Current, lines []*PickLine, err error) {
	if c, err = s.CurrentWeek(); err != nil {
		return
	}
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
			awayTeam.team_id,
			awayTeam.team_city,
			awayTeam.team_name,
			awayTeam.team_league,
			awayTeam.team_division,
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
		WHERE
			game_week     = $1
			AND game_year = $2
		ORDER BY game_start ASC, team_id_home ASC
	`
	rows, err := s.db.Query(query, c.Week, c.Year)
	if err != nil {
		return
	}
	defer rows.Close()

	lines = make([]*PickLine, 0, 16)
	for rows.Next() {
		var game_id string
		l := new(PickLine)
		err = rows.Scan(
			&l.Start,
			&game_id,
			&l.Line.Spread,
			&l.Line.OverUnder,
			&l.Line.Updated,
			&l.Home.Id,
			&l.Home.City,
			&l.Home.Name,
			&l.Home.League,
			&l.Home.Division,
			&l.Away.Id,
			&l.Away.City,
			&l.Away.Name,
			&l.Away.League,
			&l.Away.Division,
			&l.Stadium.Id,
			&l.Stadium.Name,
			&l.Stadium.City,
			&l.Stadium.State,
			&l.Stadium.Turf,
			&l.Stadium.Roof,
		)
		l.Line.GameId = GameIdType(game_id)
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
