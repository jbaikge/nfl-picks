package main

import (
	"github.com/jbaikge/nfl-picks/external/oddsmaker"
	"github.com/jbaikge/nfl-picks/picks"
)

type ImportOdds struct{}

var _ CommandInterface = new(ImportOdds)

func init() {
	RegisterCommand(&Command{
		Operation: "importodds",
		Arguments: "",
		HelpText:  "Imports odds information from OddsMaker.ag",
		Execute:   new(ImportOdds),
	})
}

func (io *ImportOdds) Do(s *picks.Store, args []string) (err error) {
	odds, err := oddsmaker.CurrentOdds()
	if err != nil {
		return
	}
	for _, o := range odds {
		odd := &picks.Odds{
			GameId:    picks.GameId(o.Home.Name, o.Away.Name, o.GameTime),
			Spread:    o.Home.Spread,
			OverUnder: o.Home.Total,
		}
		if err = s.SaveOdds(odd); err != nil {
			return
		}
	}
	return
}
