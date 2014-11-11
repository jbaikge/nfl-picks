package picks

import (
	"fmt"
	"time"
)

type Quarter string

type Game struct {
	Id        string
	EventId   int64
	Year      int
	Week      int
	Start     time.Time
	TimeLeft  time.Duration
	Posession string
	HomeId    string
	HomeScore int
	AwayId    string
	AwayScore int
	Quarter   Quarter
}

const (
	Pregame       Quarter = "P"
	FirstQuarter          = "1"
	SecondQuarter         = "2"
	Halftime              = "H"
	ThirdQuarter          = "3"
	FourthQuarter         = "4"
	Final                 = "F"
	FinalOvertime         = "FO"
)

func (q Quarter) String() string {
	switch q {
	case Pregame:
		return "Pregame"
	case FirstQuarter:
		return "1st"
	case SecondQuarter:
		return "2nd"
	case Halftime:
		return "Halftime"
	case ThirdQuarter:
		return "3rd"
	case FourthQuarter:
		return "4th"
	case Final:
		return "Final"
	case FinalOvertime:
		return "Final (Overtime)"
	}
	return "UNKOWN [" + string(q) + "]"
}

func GameId(home, away string, date time.Time) (id string) {
	return fmt.Sprintf("%sv%s@%s", home, away, date.Format("20060102"))
}
