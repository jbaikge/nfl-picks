package main

import (
	"github.com/jbaikge/nfl-picks/bin/picks-web/apitypes"
)

type Picks struct{}

func init() {
	RegisterAPI(new(Picks))
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
	return
}
