package main

import (
	"database/sql"
	"log"

	"github.com/davisbento/simple-bank-golang/api"
	db "github.com/davisbento/simple-bank-golang/db/sqlc"
	"github.com/davisbento/simple-bank-golang/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config file: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
