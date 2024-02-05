package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	api_v1 "github.com/kimhieu153255/first-go/internal/api/v1"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
	config_env "github.com/kimhieu153255/first-go/internal/config/env"
)

func main() {
	config, err := config_env.NewConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		panic(err)
	}

	store := db.NewStore(conn)

	server, err := api_v1.NewServer(store, config)
	if err != nil {
		panic(err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
