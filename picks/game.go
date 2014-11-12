package picks

import (
	"fmt"
	"time"
)

type Game struct {
	Id        string
	EventId   int64
	Year      int
	Week      int
	Start     time.Time
	TimeLeft  time.Duration
	Posession string
	Home      string
	HomeScore int
	Away      string
	AwayScore int
	Quarter   Quarter
	Line      Line
}

type Line struct {
	Spread    float64
	OverUnder float64
	Updated   time.Time
}

func GameId(away, home string, date time.Time) (id string) {
	return fmt.Sprintf("%sv%s@%s", away, home, date.Format("20060102"))
}
