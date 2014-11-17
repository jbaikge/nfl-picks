package main

import (
	"fmt"
	"github.com/jbaikge/nfl-picks/bin/picks-web/apitypes"
)

type User struct{}

func init() {
	RegisterAPI(new(User))
}

func (api *User) Auth(in *apitypes.UserAuthIn, out *apitypes.UserAuthOut) (err error) {
	out.Id, out.IsAdmin, err = Store.UserValidate(in.Username, in.PIN)
	out.Username = in.Username
	return
}

func (api *User) LastSeen(in *apitypes.UserLastSeenIn, out *Nil) (err error) {
	err = Store.UserLastSeen(in.Id)
	return
}

func (api *User) Update(in *apitypes.UserUpdateIn, out *Nil) (err error) {
	match, err := Store.UserValidatePIN(in.Id, in.OldPIN)
	if err != nil {
		return
	}
	if !match {
		err = fmt.Errorf("Incorrect PIN")
	}
	err = Store.UpdateUser(in.Id, in.NewUsername, in.NewPIN)
	return
}

func (api *User) Usernames(in *Nil, out *apitypes.UserUsernamesOut) (err error) {
	out.Usernames, err = Store.Usernames()
	return
}
