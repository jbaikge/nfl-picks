package main

import (
	"github.com/jbaikge/nfl-picks/bin/picks-web/apitypes"
)

type Picks struct{}

func init() {
	RegisterAPI(new(Picks))
}

func (p *Picks) AllCurrent(in *Nil, out *apitypes.PicksAllOut) (err error) {
	out.Week, err = Store.CurrentWeek()
	if err != nil {
		return
	}
	out.Picks, err = Store.AllPicks(out.Week)
	if err != nil {
		return
	}
	return
}

func (p *Picks) Submit(in *apitypes.PicksSubmitIn, out *apitypes.PicksSubmitOut) (err error) {
	out.Valid = make([]bool, len(in.Picks))
	out.Saved = true
	for i, pick := range in.Picks {
		valid := pick.Valid()
		out.Valid[i] = valid
		out.Saved = out.Saved && valid
	}
	if !out.Saved {
		return
	}
	for _, pick := range in.Picks {
		if err = Store.Pick(in.UserId, &pick); err != nil {
			return
		}
	}
	return
}
