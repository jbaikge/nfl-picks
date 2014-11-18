package picks

import (
	"fmt"
	"strings"
	"time"
)

type GameIdType string

type Game struct {
	Id        GameIdType
	EventId   int64
	Year      int
	Week      int
	Season    string
	Start     time.Time
	TimeLeft  time.Duration
	Posession string
	Home      string
	HomeScore int
	Away      string
	AwayScore int
	Quarter   Quarter
}

func GameId(away, home string, date time.Time) (id GameIdType) {
	return GameIdType(fmt.Sprintf("%sv%s@%s", away, home, date.Format("20060102")))
}

func (g GameIdType) Away() string {
	s := g.String()
	return s[:strings.Index(s, "v")]
}

func (g GameIdType) Home() string {
	s := g.String()
	return s[strings.Index(s, "v")+1 : strings.Index(s, "@")]
}

func (g GameIdType) String() string {
	return string(g)
}
