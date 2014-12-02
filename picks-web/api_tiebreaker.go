package main

import (
	"github.com/jbaikge/nfl-picks/picks"
)

type TieBreaker struct{}

func init() {
	RegisterAPI(new(TieBreaker))
}

// All / Current

type AllTieBreakerIn struct {
	Week picks.Week
}

type AllTieBreakerOut struct {
	TieBreakers map[string]string
}

func (api *TieBreaker) All(in *AllTieBreakerIn, out *AllTieBreakerOut) (err error) {

	return
}

func (api *TieBreaker) Current(in *Nil, out *AllTieBreakerOut) (err error) {
	currentIn := new(AllTieBreakerIn)
	if currentIn.Week, err = Store.CurrentWeek(); err != nil {
		return
	}
	return api.All(currentIn, out)
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
