package main

import (
	"fmt"
	"github.com/jbaikge/nfl-picks/picks"
)

type Help struct{}

var _ CommandInterface = new(Help)

func init() {
	RegisterCommand(&Command{
		Operation: "help",
		Arguments: "",
		HelpText:  "This help information",
		Execute:   new(Help),
		SkipDB:    true,
	})
}

func (_ *Help) Do(s *picks.Store, args []string) (err error) {
	fmt.Println("Available commands")
	var cmdWidth = 0
	for _, cmd := range commands {
		if w := len(cmd.Operation) + len(cmd.Arguments); w > cmdWidth {
			cmdWidth = w
		}
	}
	layout := fmt.Sprintf("%%-%ds %%s\n", cmdWidth+1)
	for _, cmd := range commands {
		op := cmd.Operation
		if cmd.Arguments != "" {
			op += " " + cmd.Arguments
		}
		fmt.Printf(layout, op, cmd.HelpText)
	}
	return
}
