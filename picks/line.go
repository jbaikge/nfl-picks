package picks

import (
	"time"
)

type Line struct {
	GameId    GameIdType
	Spread    float64
	OverUnder float64
	Updated   time.Time
}
