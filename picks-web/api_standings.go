package main

import (
	"github.com/jbaikge/nfl-picks/picks"
)

type Standings struct{}

func init() {
	RegisterAPI(new(Standings))
}

// Overall

type OverallOut struct {
	Standings []picks.WeekStanding
}

func (api *Standings) Overall(in *Nil, out *OverallOut) (err error) {
	out.Standings, err = Store.Standings(2014)
	return
}
