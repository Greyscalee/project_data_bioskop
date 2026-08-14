package main

import (
	"data-bioskop/routers"
	"database/sql"
	"fmt"

	_"github.com/lib/pq"
)


const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "inlinesk8"
	dbname = "postgres"
)

var (
	db *sql.DB
	err error
)

func main() {
	var PORT = ":8080"
	routers.StartServer().Run(PORT)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

