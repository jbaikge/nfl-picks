package main

import (
	"fmt"
	"github.com/jbaikge/nfl-picks/external/nfl"
	"github.com/jbaikge/nfl-picks/picks"
	"strconv"
)

type UpdateScores struct{}

var _ CommandInterface = new(UpdateScores)

func init() {
	RegisterCommand(&Command{
		Operation: "updatescores",
		Arguments: "",
		HelpText:  "Updates score information with current data",
		Execute:   new(UpdateScores),
	})
}

func (us *UpdateScores) Do(s *picks.Store, args []string) (err error) {
	year, week, games, err := nfl.CurrentGames()
	if err != nil {
		return
	}
	for _, game := range games {
		if q := picks.Quarter(game.Quarter); q == picks.Pregame {
			fmt.Printf("Skipping %s during %s\n", game.Id, q)
			continue
		}
		if err = s.SaveGame(us.convert(game, year, week)); err != nil {
			return
		}
		fmt.Printf("Updated score %s %d %s %d\n", game.Away, game.AwayScore, game.Home, game.HomeScore)
	}
	return
}
