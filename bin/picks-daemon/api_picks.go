package main

type Picks struct {}

func init(){
	RegisterAPI(new(Picks))
}

func (p *Picks) Submit(in *Nil, out *Nil) (err error) {
	return
}
