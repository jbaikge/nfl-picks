package main

import (
	"flag"
	"fmt"
	"github.com/jbaikge/nfl-picks/picks"
	"log"
	"os"
)

var dsn = flag.String("dsn", os.Getenv("DATABASE_URL"), "Data Source Name")

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = append(args, "help")
	}

	cmd, ok := commands[flag.Arg(0)]
	if !ok {
		fmt.Printf("Unknown command: %s\n", flag.Arg(0))
		cmd = commands["help"]
	}

	var s *picks.Store
	if !cmd.SkipDB {
		var err error
		s, err = picks.NewStore(*dsn)
		if err != nil {
			log.Fatalf("NewStore: %s", err)
		}
		defer s.Close()
	}

	if err := cmd.Execute.Do(s, args[1:]); err != nil {
		fmt.Printf("There was an error while processing %s:\n%s\n", cmd.Operation, err)
		os.Exit(1)
	}
}
