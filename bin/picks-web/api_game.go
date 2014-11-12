package main

import (
	"fmt"
	"github.com/jbaikge/nfl-picks/external/nfl"
	"github.com/jbaikge/nfl-picks/picks"
)

type Game struct{}

type GameImportIn struct {
	Week int
	Year int
}

type GameImportOut struct {
	Games []*picks.Game
}

func init() {
	RegisterAPI(new(Game))
}

func (api *Game) ImportWeek(in *GameImportIn, out *GameImportOut) (err error) {
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

func (api *Game) ImportYear(in *GameImportIn, out *GameImportOut) (err error) {
	out.Games = make([]*picks.Game, 0, 300)
	for week := 1; week <= 17; week++ {
		req := &GameImportIn{
			Week: week,
			Year: in.Year,
		}
		resp := new(GameImportOut)
		if err = api.ImportWeek(req, resp); err != nil {
			return
		}
		out.Games = append(out.Games, resp.Games...)
	}
	return
}
