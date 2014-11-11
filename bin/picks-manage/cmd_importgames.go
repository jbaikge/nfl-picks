package main

import (
	"fmt"
	"github.com/jbaikge/nfl-picks/external/nfl"
	"github.com/jbaikge/nfl-picks/picks"
	"strconv"
)

type ImportGames struct{}

var _ CommandInterface = new(ImportGames)

func init() {
	RegisterCommand(&Command{
		Operation: "importgames",
		Arguments: "year [week]",
		HelpText:  "Imports game information from NFL.com. No week arg will import entire year",
		Execute:   new(ImportGames),
	})
}

func (ig *ImportGames) Do(s *picks.Store, args []string) (err error) {
	if a := len(args); a < 1 || a > 2 {
		return fmt.Errorf("Expected 1 or 2 arguments, got %d", a)
	}
	var weekGames []nfl.Game
	var year, week int

	if year, err = strconv.Atoi(args[0]); err != nil {
		return fmt.Errorf("Atoi(%d): %s", year, err)
	}

	switch len(args) {
	case 1:
		for week = 1; week <= 17; week++ {
			if weekGames, err = nfl.GamesFor(year, week); err != nil {
				return
			}
			for _, g := range weekGames {
				if err = s.SaveGame(ig.convert(g, year, week)); err != nil {
					return
				}
			}
		}
	case 2:
		if week, err = strconv.Atoi(args[1]); err != nil {
			return fmt.Errorf("Atoi(%d): %s", week, err)
		}
		if weekGames, err = nfl.GamesFor(year, week); err != nil {
			return
		}
		for _, g := range weekGames {
			if err = s.SaveGame(ig.convert(g, year, week)); err != nil {
				return
			}
		}
	}

	return
}

func (ig *ImportGames) convert(in nfl.Game, year, week int) (out *picks.Game) {
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
