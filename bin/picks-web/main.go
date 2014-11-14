package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jbaikge/nfl-picks/picks"
	"log"
	"net/http"
	"os"
)

var (
	Port      = flag.String("port", os.Getenv("PORT"), "Listen port")
	DSN       = flag.String("dsn", os.Getenv("DATABASE_URL"), "Data Source Name")
	AssetsDir = flag.String("assets", "./assets", "Assets directory")
	SetupDB   = flag.Bool("setupdb", false, "Update database structure")
)

var (
	Store  *picks.Store
	Router = mux.NewRouter()
)

func main() {
	flag.Parse()

	// Database
	var err error
	Store, err = picks.NewStore(*DSN)
	if err != nil {
		log.Fatalf("NewStore: %s", err)
	}
	defer Store.Close()

	if *SetupDB {
		if err := Store.Setup(); err != nil {
			log.Fatalf("Store.Setup: %s", err)
		}
	}

	// Routing
	Router.PathPrefix("/").Handler(http.FileServer(http.Dir(*AssetsDir)))
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, Router))

	log.Fatal(http.ListenAndServe(":"+*Port, nil))
}
