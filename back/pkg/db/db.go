package db

import (
	"database/sql"
	"extractor/back/pkg/env"
	"extractor/back/pkg/er"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func Connect() (db *sql.DB) {
	dbDriver := env.Get("DB_DRIVER")
	dbServer := env.Get("DB_SERVER")
	dbUser := env.Get("DB_USER")
	dbPassword := env.Get("DB_PASSWORD")
	dbName := env.Get("DB_NAME")

	db, err := sql.Open(dbDriver, "server="+dbServer+";user id="+dbUser+";password="+dbPassword+";database="+dbName)

	er.Check(err)

	return db
}

func InvokeQuery(query string, args ...interface{}) *sql.Rows {
	db := Connect()
	fmt.Println(query)
	rows, err := db.Query(query, args...)
	defer db.Close()
	er.Check(err)
	return rows
}

func InvokeQueryRow(query string, args ...interface{}) *sql.Row {
	db := Connect()
	row := db.QueryRow(query, args...)
	defer db.Close()
	return row
}

func InvokeExec(query string, args ...interface{}) sql.Result {
	db := Connect()
	result, err := db.Exec(query, args...)
	defer db.Close()
	er.Check(err)
	return result
}
