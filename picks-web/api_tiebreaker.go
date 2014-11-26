package main

import (
	"github.com/jbaikge/nfl-picks/picks"
)

type TieBreaker struct{}

func init() {
	RegisterAPI(new(TieBreaker))
}

// Submit

type SubmitTieBreakerIn struct {
	UserId     int64
	TieBreaker picks.TieBreaker
}

type SubmitTieBreakerOut struct {
	Saved bool
}

func (api *TieBreaker) Submit(in *SubmitTieBreakerIn, out *SubmitTieBreakerOut) (err error) {
	if err = Store.AddTieBreaker(in.UserId, in.TieBreaker); err != nil {
		return
	}
	out.Saved = true
	return
}
