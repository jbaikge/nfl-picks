package main

import (
	"net/http"
)

type User struct{}

type UserAuthIn struct {
	Name  string
	Phone string
}

type UserAuthOut struct {
	Valid bool
}

func init() {
	RegisterAPI(new(User))
}

func (u *User) Auth(r *http.Request, in *UserAuthIn, out *UserAUthOut) (err error) {
	return
}
