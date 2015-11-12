package picks

type Team struct {
	Id       string
	Name     string
	City     string
	League   string
	Division string
	Wins     int
	Losses   int
}

type TeamScore struct {
	Team  Team
	Score int
}
