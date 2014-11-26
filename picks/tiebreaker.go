package picks

type TieBreaker struct {
	GameId GameIdType
	Value  string
}

func (t TieBreaker) Valid() bool {
	switch t.Value {
	case "ROCK", "PAPER", "SCISSORS", "LIZARD", "SPOCK":
		return true
	}
	return false
}
