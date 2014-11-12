package main

import (
	"flag"
	"github.com/jbaikge/nfl-picks/picks"
	"log"
	"net/http"
	"os"
)

var (
	ListenAddr = flag.String("listen", ":10000", "Listen address")
	DSN        = flag.String("dsn", os.Getenv("DATABASE_URL"), "Data Source Name")
)

var (
	Store *picks.Store
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = append(args, "help")
	}

	var err error
	Store, err = picks.NewStore(*DSN)
	if err != nil {
		log.Fatalf("NewStore: %s", err)
	}
	defer Store.Close()

	log.Fatal(http.ListenAndServe(*ListenAddr, nil))
}
