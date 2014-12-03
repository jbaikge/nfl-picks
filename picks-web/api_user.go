package main

import (
	"fmt"
)

type User struct{}

func init() {
	RegisterAPI(new(User))
}

// Auth

type AuthIn struct {
	Username string
	PIN      int
}

type AuthOut struct {
	Id       int64
	IsAdmin  bool
	Username string
	Theme    string
	Beer     string
}

func (api *User) Auth(in *AuthIn, out *AuthOut) (err error) {
	out.Id, out.IsAdmin, out.Theme, out.Beer, err = Store.UserValidate(in.Username, in.PIN)
	out.Username = in.Username
	return
}

// Last Seen

type LastSeenIn struct {
	Id int64
}

func (api *User) LastSeen(in *LastSeenIn, out *Nil) (err error) {
	err = Store.UserLastSeen(in.Id)
	return
}

// Update

type UpdateIn struct {
	Id          int64
	NewUsername string
	Beer        string
	OldPIN      int
	NewPIN      int
	Theme       string
}

type UpdateOut struct {
	Id       int64
	IsAdmin  bool
	Theme    string
	Username string
}

func (api *User) Update(in *UpdateIn, out *Nil) (err error) {
	match, err := Store.UserValidatePIN(in.Id, in.OldPIN)
	if err != nil {
		return
	}
	if !match {
		err = fmt.Errorf("Incorrect PIN")
		return
	}
	err = Store.UpdateUser(in.Id, in.NewUsername, in.Beer, in.NewPIN, in.Theme)
	return
}

// Usernames

type UsernamesOut struct {
	Usernames []string
}

func (api *User) Usernames(in *Nil, out *UsernamesOut) (err error) {
	out.Usernames, err = Store.Usernames()
	return
}
