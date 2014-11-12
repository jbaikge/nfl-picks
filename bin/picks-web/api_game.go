package main

type Game struct{}

func init() {
	RegisterAPI(new(Game))
}

func (g *Game) ImportWeek(in *Nil, out *Nil) (err error) {
	return
}

func (g *Game) ImportYear(in *Nil, out *Nil) (err error) {
	return
}
