package tests

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
	config_env "github.com/kimhieu153255/first-go/internal/config/env"
	_ "github.com/lib/pq"
)

var testStore db.Store

func TestMain(m *testing.M) {
	config, err := config_env.NewConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = db.NewStore(connPool)
	os.Exit(m.Run())
}
