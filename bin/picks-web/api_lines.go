package main

import (
	"fmt"
	"github.com/jbaikge/nfl-picks/bin/picks-web/apitypes"
	"github.com/jbaikge/nfl-picks/external/oddsmaker"
	"github.com/jbaikge/nfl-picks/picks"
)

type Lines struct{}

func init() {
	RegisterAPI(new(Lines))
}

func (api *Lines) Current(in *apitypes.LinesCurrentIn, out *apitypes.LinesCurrentOut) (err error) {
	out.Current, out.Lines, err = Store.CurrentPickLines()
	if err != nil {
		return
	}
	if in.UserId > 0 {
		out.Picks, err = Store.UserPicks(in.UserId, out.Current)
	}
	return
}

func (api *Lines) ImportCurrent(in *Nil, out *apitypes.LinesImportOut) (err error) {
	out.Lines, err = oddsmaker.CurrentLines()
	if err != nil {
		return
	}
	for _, line := range out.Lines {
		fmt.Printf("api.Lines.ImportCurrent: %16s %5.1f %5.1f\n", line.GameId, line.Spread, line.OverUnder)
		if err = Store.UpdateLine(line); err != nil {
			return
		}
	}
	return
}

func (api *Lines) Backfill(in *picks.Line, out *picks.Line) (err error) {
	fmt.Printf("api.Lines.Backfill: %16s %5.1f %5.1f\n", in.GameId, in.Spread, in.OverUnder)
	if err = Store.UpdateLine(in); err != nil {
		return
	}
	*out = *in
	return
}
