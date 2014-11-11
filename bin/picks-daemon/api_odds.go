package main

type Odds struct {}

func init() {
	RegisterAPI(new(Odds))
}

func (o *Odds) ImportCurrent(in *Nil, out *Nil) (err error) {
	return
}
