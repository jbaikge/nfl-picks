package main

type User struct{}

type UserAuthOut struct {
	Valid bool
}

func init() {
	RegisterAPI(new(User))
}

func (api *User) Auth(in *Nil, out *Nil) (err error) {
	return
}
