package main

type User struct{}

func init(){
	RegisterAPI(new(User))
}

func (u *User) Auth(in *Nil, out *Nil) (err error) {
	return
}
