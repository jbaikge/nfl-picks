package picks

type Team struct {
	Id   string
	Name string
	City string
}

type TeamScore struct {
	Team  Team
	Score int
}
