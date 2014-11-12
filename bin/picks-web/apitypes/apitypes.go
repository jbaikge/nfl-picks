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

type LinesImportOut struct {
	Lines []*picks.Line
}
