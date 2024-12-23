package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/todevmilen/simplebank/api"
	db "github.com/todevmilen/simplebank/db/sqlc"
	"github.com/todevmilen/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
