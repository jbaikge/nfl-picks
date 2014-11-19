package apitypes

import (
	"github.com/jbaikge/nfl-picks/picks"
)

type GameCurrentGamesOut struct {
	Current picks.Current
	Games   []*picks.Game
}

type GameCurrentWeekOut struct {
	picks.Current
}

type GameImportIn struct {
	Year int
	Week int
}

type GameImportOut struct {
	Games []*picks.Game
}

type GameScoresOut struct {
	Updated []*picks.Game
}

type LinesCurrentIn struct {
	UserId int64
}

type LinesCurrentOut struct {
	Current picks.Current
	Lines   []*picks.PickLine
	Picks   []*picks.Pick
}

type LinesImportOut struct {
	Lines []*picks.Line
}

type PicksAllOut struct {
	Current picks.Current
	Picks   map[picks.GameIdType]map[string]picks.Pick
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

type UserLastSeenIn struct {
	Id int64
}

type UserUpdateIn struct {
	Id          int64
	NewUsername string
	OldPIN      int
	NewPIN      int
}

type UserUpdateOut struct {
	Id       int64
	IsAdmin  bool
	Username string
}

type UserUsernamesOut struct {
	Usernames []string
}
