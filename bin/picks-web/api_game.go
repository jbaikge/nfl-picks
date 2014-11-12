package main

import (
	"fmt"
	"github.com/jbaikge/nfl-picks/bin/picks-web/apitypes"
	"github.com/jbaikge/nfl-picks/external/nfl"
	"github.com/jbaikge/nfl-picks/picks"
)

type Game struct{}

func init() {
	RegisterAPI(new(Game))
}

func (api *Game) ImportWeek(in *apitypes.GameImportIn, out *apitypes.GameImportOut) (err error) {
	if out.Games, err = nfl.GamesFor(in.Year, in.Week); err != nil {
		return
	}
	for _, g := range out.Games {
		fmt.Printf("api.Game.ImportWeek: %s\n", g.Id)
		if err = Store.NewGame(g); err != nil {
			return
		}
	}
	return
}

func (api *Game) ImportYear(in *apitypes.GameImportIn, out *apitypes.GameImportOut) (err error) {
	out.Games = make([]*picks.Game, 0, 300)
	for week := 1; week <= 17; week++ {
		req := &apitypes.GameImportIn{
			Week: week,
			Year: in.Year,
		}
		resp := new(apitypes.GameImportOut)
		if err = api.ImportWeek(req, resp); err != nil {
			return
		}
		out.Games = append(out.Games, resp.Games...)
	}
	return
}
