package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kennyclark/web-pos-api/internal/db/sqlc"
)

var pool *pgxpool.Pool
var Q *sqlc.Queries

func Connect(dbURL string) {
	log.Println("Connecting to database...")
	var err error
	pool, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	if err = pool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	log.Println("Connected to database.")
	Q = sqlc.New(pool)
}

func Disconnect() {
	if pool != nil {
		log.Println("Disconnecting from database...")
		pool.Close()
		log.Println("Disconnected from database")
	}
}
