package picks

import (
	"database/sql"
)

type WeekStanding struct {
	Week       Week
	UserWins   map[string]int
	Winners    []string
	TotalGames int
}

func (s *Store) Standings(year int) (standings []WeekStanding, err error) {
	standings = make([]WeekStanding, 0, 17)

	var rows *sql.Rows

	userWins := make(map[int]map[string]int)
	queryWins := `SELECT
			games.game_week,
			users.user_name,
			COUNT(0) AS correct
		FROM
			picks
			LEFT JOIN games USING(game_id)
			LEFT JOIN users USING(user_id)
		WHERE
			games.game_year = $1
			AND (
				(
					picks.pick_value = games.team_id_away
					AND games.game_score_away - games.game_spread >= games.game_score_home
				) OR (
					picks.pick_value = games.team_id_home
					AND games.game_score_home + games.game_spread >= games.game_score_away
				) OR (
					picks.pick_value = 'UNDER'
					AND games.game_score_away + games.game_score_home <= games.game_over_under
				) OR (
					picks.pick_value = 'OVER'
					AND games.game_score_away + games.game_score_home >= games.game_over_under
				)
			)
		GROUP BY
			users.user_id,
			games.game_week
	`
	if rows, err = s.db.Query(queryWins, year); err != nil {
		return
	}
	for rows.Next() {
		var week, correct int
		var user string
		if err = rows.Scan(&week, &user, &correct); err != nil {
			return
		}
		if userWins[week] == nil {
			userWins[week] = make(map[string]int)
		}
		userWins[week][user] = correct
	}
	if err = rows.Close(); err != nil {
		return
	}

	// Determine winners, hit up tie-breaker if necessary

	queryWeeks := `SELECT game_week, COUNT(*) FROM games WHERE game_year = $1 GROUP BY game_week ORDER BY game_week`
	if rows, err = s.db.Query(queryWeeks, year); err != nil {
		return
	}
	for rows.Next() {
		var week, count int
		if err = rows.Scan(&week, &count); err != nil {
			return
		}
		ws := WeekStanding{
			Week: Week{
				Week:   week,
				Year:   year,
				Season: "R",
			},
			TotalGames: count,
			Winners:    make([]string, 0, 8),
			UserWins:   userWins[week],
		}
		standings = append(standings, ws)
	}
	if err = rows.Close(); err != nil {
		return
	}

	return
}
