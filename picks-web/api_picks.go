package main

import (
	"github.com/jbaikge/nfl-picks/picks"
	"time"
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
	Week   picks.Week
	Picks  map[picks.GameIdType]map[string]picks.Pick
	Totals map[string]int
	Winner string
}

func (api *Picks) All(in *AllIn, out *AllOut) (err error) {
	if out.Picks, err = Store.AllPicks(in.Week); err != nil {
		return
	}
	out.Totals = make(map[string]int)
	for id := range out.Picks {
		for username, pick := range out.Picks[id] {
			if pick.Correct {
				out.Totals[username]++
			}
		}
	}
	out.Week = in.Week
	return
}

// All Current

func (api *Picks) AllCurrent(in *Nil, out *AllOut) (err error) {
	if out.Week, err = Store.CurrentWeek(); err != nil {
		return
	}
	return api.All(&AllIn{Week: out.Week}, out)
}

// Closed

type ClosedOut struct {
	Closed bool
}

func (api *Picks) Closed(in *Nil, out *ClosedOut) (err error) {
	w, err := Store.CurrentWeek()
	if err != nil {
		return
	}
	t, err := Store.FirstGameTime(w)
	if err != nil {
		return
	}
	out.Closed = time.Now().After(t)
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

func (api *Picks) Submit(in *SubmitPickIn, out *SubmitPickOut) (err error) {
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
