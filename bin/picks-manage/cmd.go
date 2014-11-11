package main

import (
	"github.com/jbaikge/nfl-picks/picks"
	"log"
)

type Command struct {
	Operation string
	Arguments string
	HelpText  string
	Execute   CommandInterface
	SkipDB    bool
}

type CommandInterface interface {
	Do(s *picks.Store, args []string) (err error)
}

var commands = make(map[string]*Command, 10)

func RegisterCommand(c *Command) {
	if _, ok := commands[c.Operation]; ok {
		log.Fatalf("Command for operation `%s' already exists", c.Operation)

	}
	commands[c.Operation] = c
}
