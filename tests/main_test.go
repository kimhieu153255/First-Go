package tests

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/kimhieu/first-go/internal/config/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "pgx"
	dbSource = "postgres://root:secret@localhost:5432/testGo?sslmode=disable"
)

var testStore db.Store

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = db.NewStore(connPool)
	os.Exit(m.Run())
}
