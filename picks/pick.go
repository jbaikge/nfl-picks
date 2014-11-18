package picks

import (
	"time"
)

type Pick struct {
	GameId GameIdType
	Value  string
}

type PickLine struct {
	Line    Line
	Start   time.Time
	Home    Team
	Away    Team
	Stadium Stadium
}

type Stadium struct {
	Id    string
	Name  string
	City  string
	State string
	Turf  string
	Roof  string
}

func (p Pick) Valid() bool {
	switch p.Value {
	case "OVER", "UNDER", p.GameId.Away(), p.GameId.Home():
		return true
	default:
		return false
	}
}
