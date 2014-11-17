package apitypes

import (
	"github.com/jbaikge/nfl-picks/picks"
)

type GameImportIn struct {
	Week int
	Year int
}

type GameImportOut struct {
	Games []*picks.Game
}

type GameScoresOut struct {
	Updated []*picks.Game
}

type LinesImportOut struct {
	Lines []*picks.Line
}

type PicksSubmitIn struct {
	UserId int64
	Picks  []picks.Pick
}

type PicksSubmitOut struct {
	Valid []bool
	Saved bool
}

type UserAuthIn struct {
	Username string
	PIN      int
}

type UserAuthOut struct {
	Id       int64
	IsAdmin  bool
	Username string
}

type UserUsernamesOut struct {
	Usernames []string
}
