package main

import (
	"github.com/jbaikge/nfl-picks/bin/picks-web/apitypes"
	"log"
)

type User struct{}

func init() {
	RegisterAPI(new(User))
}

func (api *User) Auth(in *apitypes.UserAuthIn, out *apitypes.UserAuthOut) (err error) {
	log.Printf("Username: `%s', PIN: `%s'", in.Username, in.PIN)
	out.Id, out.IsAdmin, err = Store.UserValidate(in.Username, in.PIN)
	out.Username = in.Username
	return
}

func (api *User) Usernames(in *Nil, out *apitypes.UserUsernamesOut) (err error) {
	out.Usernames, err = Store.Usernames()
	return
}
