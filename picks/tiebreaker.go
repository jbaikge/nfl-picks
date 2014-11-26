package picks

type TieBreaker struct {
	Week  Week
	Value string
}

func (t TieBreaker) Valid() bool {
	switch t.Value {
	case "ROCK", "PAPER", "SCISSORS", "LIZARD", "SPOCK":
		return true
	}
	return false
}
