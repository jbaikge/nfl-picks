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
		if q := picks.Quarter(game.Quarter); q == picks.Pregame || q == picks.Final || q == picks.FinalOvertime {
			continue
		}
		if err = s.SaveGame(us.convert(game, year, week)); err != nil {
			return
		}
		fmt.Printf("Updated score %s %d %s %d\n", game.Away, game.AwayScore, game.Home, game.HomeScore)
	}
	return
}

func (us *UpdateScores) convert(in nfl.Game, year, week int) (out *picks.Game) {
	out = &picks.Game{
		Week:      week,
		Year:      year,
		Start:     in.Start(),
		TimeLeft:  in.TimeLeft(),
		Posession: in.Posession,
		HomeId:    in.Home,
		HomeScore: in.HomeScore,
		AwayId:    in.Away,
		AwayScore: in.AwayScore,
		Quarter:   picks.Quarter(in.Quarter),
	}
	out.EventId, _ = strconv.ParseInt(in.EventId, 10, 64)
	return
}
