package main

import (
	"fmt"
	"github.com/jbaikge/nfl-picks/external/nfl"
	"github.com/jbaikge/nfl-picks/picks"
	"time"
)

type Game struct{}

func init() {
	RegisterAPI(new(Game))
}

// Current Week

type CurrentWeekOut struct {
	picks.Week
}

func (api *Game) CurrentWeek(in *Nil, out *CurrentWeekOut) (err error) {
	out.Week, err = Store.CurrentWeek()
	return
}

// Update Current Week

func (api *Game) UpdateCurrentWeek(in *Nil, out *CurrentWeekOut) (err error) {
	w, _, err := nfl.CurrentGames()
	if err != nil {
		return
	}
	if err = Store.UpdateCurrentWeek(w.Year, w.Week, w.Season); err != nil {
		return
	}
	out.Week = w
	return
}

// Import Week

type ImportIn struct {
	Year int
	Week int // Optional for ImportYear
}

type ImportOut struct {
	Games []*picks.Game
}

func (api *Game) ImportWeek(in *ImportIn, out *ImportOut) (err error) {
	if out.Games, err = nfl.GamesFor(in.Year, in.Week); err != nil {
		return
	}
	for _, g := range out.Games {
		fmt.Printf("api.Game.ImportWeek: %16s\n", g.Id)
		if err = Store.NewGame(g); err != nil {
			return
		}
	}
	return
}

// Import Year

func (api *Game) ImportYear(in *ImportIn, out *ImportOut) (err error) {
	out.Games = make([]*picks.Game, 0, 300)
	for week := 1; week <= 17; week++ {
		req := &ImportIn{
			Week: week,
			Year: in.Year,
		}
		resp := new(ImportOut)
		if err = api.ImportWeek(req, resp); err != nil {
			return
		}
		out.Games = append(out.Games, resp.Games...)
	}
	return
}

// Scores

type ScoresOut struct {
	NextUpdate time.Duration
	Scores     []*picks.Game
}

func (api *Game) Scores(in *picks.Week, out *ScoresOut) (err error) {
	out.Scores, err = Store.Scores(*in)
	if err != nil {
		return
	}
	out.NextUpdate = time.Hour
	for _, s := range out.Scores {
		// Skip completed games
		if s.Quarter == picks.Final || s.Quarter == picks.FinalOvertime {
			continue
		}
		// Game hasn't started yet
		if s.Quarter == picks.Pregame {
			if diff := s.Start.Sub(time.Now()); diff > 0 && diff < out.NextUpdate {
				out.NextUpdate = diff
			}
			continue
		}
		// Game is in progress
		out.NextUpdate = time.Minute
	}
	return
}

// Update Scores

type UpdateScoresOut struct {
	Updated []*picks.Game
}

func (api *Game) UpdateScores(in *Nil, out *UpdateScoresOut) (err error) {
	_, games, err := nfl.CurrentGames()
	if err != nil {
		return
	}
	out.Updated = make([]*picks.Game, 0, len(games))
	for _, g := range games {
		if g.Quarter == picks.Pregame {
			continue
		}
		fmt.Printf("api.Game.UpdateScores: %16s %2d - %d\n", g.Id, g.AwayScore, g.HomeScore)
		if err = Store.UpdateGame(g); err != nil {
			return
		}
		out.Updated = append(out.Updated, g)
	}
	return
}
