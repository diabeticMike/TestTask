package datastore

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/TestTask/config"
)

// NewDB return database session
func NewDB(conf config.MySQLConfig) (*sql.DB, error) {
	mysqlInfo := fmt.Sprintf("%s:%s@(%s:%d)/%s",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName)

	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
