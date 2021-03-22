package db

import (
	"database/sql"
	"pcList/pkg/env"
	"pcList/pkg/util"

	_ "github.com/denisenkom/go-mssqldb"
)

func Connect() (db *sql.DB) {
	dbDriver := env.Get("DB_DRIVER")
	dbServer := env.Get("DB_SERVER")
	dbUser := env.Get("DB_USER")
	dbPassword := env.Get("DB_PASSWORD")
	dbName := env.Get("DB_NAME")

	db, err := sql.Open(dbDriver, "server="+dbServer+";user id="+dbUser+";password="+dbPassword+";database="+dbName)

	util.CheckErr(err)

	return db
}
