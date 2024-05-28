package main

import (
	"database/sql"
	"log"
	"task/api"
	db "task/db/sqlc"
	"task/utils"

	_ "github.com/lib/pq"
)

func main() {
	cnf, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("config файлыг уншиж чадсангүй %v", err)
	}

	conn, err := sql.Open(cnf.DBDriver, cnf.DBSource)
	if err != nil {
		log.Fatalf("DataBase Холбогдож чадсангүй %v", err)
	}
	store := db.NewStore(conn)

	server, err := api.NewServer(cnf, store)
	if err != nil {
		log.Fatalf("Сервер үүсгэж чадсангүй")
	}

	err = server.StartServer(cnf.ServerAddress)
	if err != nil {
		log.Fatalf("Сервер асааж чадсангүй")
	}
}
