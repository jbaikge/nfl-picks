package main

import (
	"github.com/jbaikge/nfl-picks/external/oddsmaker"
	"github.com/jbaikge/nfl-picks/picks"
)

type Lines struct{}

func init() {
	RegisterAPI(new(Lines))
}

// Current

type CurrentIn struct {
	UserId int64
}

type CurrentOut struct {
	Week       picks.Week
	Lines      []*picks.PickLine
	Picks      []*picks.Pick
	TieBreaker picks.TieBreaker
}

func (api *Lines) Current(in *CurrentIn, out *CurrentOut) (err error) {
	out.Week, out.Lines, err = Store.CurrentPickLines()
	if err != nil {
		return
	}
	if in.UserId > 0 {
		if out.Picks, err = Store.UserPicks(in.UserId, out.Week); err != nil {
			return
		}
		if out.TieBreaker, err = Store.UserTieBreaker(in.UserId, out.Week); err != nil {
			return
		}
	}
	return
}

// Import Current

type ImportLinesOut struct {
	Lines []*picks.Line
}

func (api *Lines) ImportCurrent(in *Nil, out *ImportLinesOut) (err error) {
	out.Lines, err = oddsmaker.CurrentLines()
	if err != nil {
		return
	}
	for _, line := range out.Lines {
		if err = Store.UpdateLine(line); err != nil {
			return
		}
	}
	return
}

// Backfill

func (api *Lines) Backfill(in *picks.Line, out *picks.Line) (err error) {
	if err = Store.UpdateLine(in); err != nil {
		return
	}
	*out = *in
	return
}
