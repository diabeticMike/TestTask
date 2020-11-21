package main

import (
	"database/sql"
	"fmt"
	baseLog "log"
	"net/http"
	"os"

	"github.com/TestTask/seeder"
	"github.com/TestTask/web/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/TestTask/config"
	"github.com/TestTask/datastore"
	"github.com/TestTask/logger"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter listen port as command agrument")
		return
	}
	listenPort := os.Args[1]
	var (
		conf config.Configuration
		log  logger.Logger
		db   *sql.DB
		err  error
	)

	// create service configuration
	if conf, err = config.New("config.json"); err != nil {
		baseLog.Fatal(err.Error())
	}
	fmt.Println(conf)

	// create service logger
	if log, err = logger.New(conf.Log); err != nil {
		baseLog.Fatal(err.Error())
	}

	// create db connection instance
	if db, err = datastore.NewDB(conf.MySQL); err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	fmt.Println(db)

	// run seeds
	if len(os.Args) == 3 && os.Args[2] == "true" {
		if err = seeder.RunSeeds(db); err != nil {
			log.Fatal(err.Error())
		}
	}

	var (
		mainRouter *mux.Router
		headers    handlers.CORSOption
		methods    handlers.CORSOption
		origins    handlers.CORSOption
	)

	// get router with CORS parameters
	mainRouter, headers, methods, origins, err = router.New(log)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Fatal(http.ListenAndServe(":"+listenPort, handlers.CORS(headers, methods, origins)(mainRouter)))
}
