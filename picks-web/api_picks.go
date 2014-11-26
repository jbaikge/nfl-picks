package main

import (
	"github.com/jbaikge/nfl-picks/picks"
)

type Picks struct{}

func init() {
	RegisterAPI(new(Picks))
}

// All

type AllIn struct {
	Week picks.Week
}

type AllOut struct {
	Week  picks.Week
	Picks map[picks.GameIdType]map[string]picks.Pick
}

func (p *Picks) All(in *AllIn, out *AllOut) (err error) {
	if out.Picks, err = Store.AllPicks(in.Week); err != nil {
		return
	}
	out.Week = in.Week
	return
}

// All Current

func (p *Picks) AllCurrent(in *Nil, out *AllOut) (err error) {
	if out.Week, err = Store.CurrentWeek(); err != nil {
		return
	}
	if out.Picks, err = Store.AllPicks(out.Week); err != nil {
		return
	}
	return
}

// Submit

type SubmitPickIn struct {
	UserId int64
	Picks  []picks.Pick
}

type SubmitPickOut struct {
	Valid []bool
	Saved bool
}

func (p *Picks) Submit(in *SubmitPickIn, out *SubmitPickOut) (err error) {
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
		if err = Store.AddPick(in.UserId, &pick); err != nil {
			return
		}
	}
	return
}
