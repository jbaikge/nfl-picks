package main

import (
	"flag"
	"fmt"
	"github.com/jbaikge/nfl-picks/picks"
	"log"
	"os"
)

var dsn = flag.String("dsn", "picks.sqlite3", "Data Source Name")

// func importGames(s *picks.Store, year, week int) (err error) {
// 	games, err := nfl.GamesFor(year, week)
// 	if err != nil {
// 		return
// 	}
// 	for _, g := range games {
// 		if err = s.SaveGame(convert(g, year, week)); err != nil {
// 			log.Printf("Game: %+v", g)
// 			return
// 		}
// 	}
// 	return
// }

// func importOdds() (err error) {
// 	odds, err := oddsmaker.CurrentOdds()
// 	if err != nil {
// 		log.Fatalf("CurrentOdds: %s", err)
// 	}
// 	for _, o := range odds {
// 		if err = s.SaveOdds(o); err != nil {
// 			log.Printf("Odds: %+v", o)
// 			return
// 		}
// 	}
// }

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

	// if y, w := *importYear, *importWeek; y > 0 && w > 0 {
	// 	if err := importGames(s, y, w); err != nil {
	// 		log.Fatalf("importGames %d %d: %s", y, w, err)
	// 	}
	// }
	// if y, w := *importYear, *importWeek; y > 0 && w == 0 {
	// 	for i := 1; i <= 17; i++ {
	// 		if err := importGames(s, y, i); err != nil {
	// 			log.Fatalf("importGames for %d: %s", y, err)
	// 		}
	// 	}
	// }
	// if *importOdds {
	// 	importOdds()
	// }

}
