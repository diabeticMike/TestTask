package main

import (
	"database/sql"
	"fmt"
	baseLog "log"
	"os"

	"github.com/TestTask/seeds"

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
	fmt.Println(os.Args[1])
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
	fmt.Println(db)

	// run seeds
	if len(os.Args) == 3 && os.Args[2] == "true" {
		if err = seeds.RunSeeds(db); err != nil {
			log.Fatal(err.Error())
		}
	}
}
