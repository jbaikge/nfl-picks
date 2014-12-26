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
	winners := make(map[int][]string)
	queryWins := `
		SELECT
			games.game_week AS week,
			users.user_name AS username,
			COUNT(0) AS correct,
			last_game.game_score_home + last_game.game_score_away AS score,
			-- tie_breakers.tie_value AS tiebr,
			-- ABS((last_game.game_score_home + last_game.game_score_away) - tie_breakers.tie_value) AS tie_score,
			rank() OVER (PARTITION BY games.game_week ORDER BY COUNT(NULLIF(users.user_can_win, FALSE)) DESC, ABS((last_game.game_score_home + last_game.game_score_away) - tie_breakers.tie_value)) AS rank
		FROM
			picks
			LEFT JOIN games USING(game_id)
			LEFT JOIN users USING(user_id)
			LEFT JOIN tie_breakers ON(
				tie_breakers.user_id = users.user_id
				AND tie_breakers.tie_year = games.game_year
				AND tie_breakers.tie_week = games.game_week
			)
			LEFT JOIN (
				SELECT
					game_year,
					game_week,
					MAX(game_start) AS game_start
				FROM
					games
				GROUP BY
					game_year,
					game_week
			) AS last_start ON (
				games.game_year = last_start.game_year
				AND games.game_week = last_start.game_week
			)
			LEFT JOIN games AS last_game ON (
				last_game.game_year = last_start.game_year
				AND last_game.game_week = last_start.game_week
				AND last_game.game_start = last_start.game_start
			)
		WHERE
			games.game_year = $1
			AND games.game_quarter IN('F', 'FO')
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
			games.game_week,
			score,
			tie_breakers.tie_value
	`
	if rows, err = s.db.Query(queryWins, year); err != nil {
		return
	}
	for rows.Next() {
		var week, correct, rank, score int
		var user string
		if err = rows.Scan(&week, &user, &correct, &score, &rank); err != nil {
			return
		}
		if userWins[week] == nil {
			userWins[week] = make(map[string]int)
			winners[week] = make([]string, 0, 8)
		}
		userWins[week][user] = correct

		if rank == 1 {
			winners[week] = append(winners[week], user)
		}
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
			Winners:    winners[week],
			UserWins:   userWins[week],
		}
		standings = append(standings, ws)
	}
	if err = rows.Close(); err != nil {
		return
	}

	return
}
