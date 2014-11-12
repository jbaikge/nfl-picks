package picks

type Pick struct {
	GameId GameIdType
	Value  string
}

func (p Pick) Valid() bool {
	switch p.Value {
	case "OVER", "UNDER", p.GameId.Away(), p.GameId.Home():
		return true
	default:
		return false
	}
}
