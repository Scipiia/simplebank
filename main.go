package main

import (
	"database/sql"
	"fmt"
	"github.com/Sciipia/simplebank/api"
	db "github.com/Sciipia/simplebank/db/sqlc"
	"github.com/Sciipia/simplebank/util"
	_ "github.com/lib/pq"
	"log"
)

//const (
//	dbDriver      = "postgres"
//	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
//	serverAddress = "0.0.0.0:8080"
//)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations", err)
	}
	fmt.Println(config)
	log.Println(config)

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
