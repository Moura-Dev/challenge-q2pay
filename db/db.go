package db

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

var (
	err  error
	Conn *sqlx.DB
)

func StartDB() (*sqlx.DB, error) {
	ctx := context.Background()
	Conn, err = sqlx.Connect("postgres", os.Getenv("DB_URL"))
	if err != nil {
		fmt.Printf("Error connecting to the database: %s", err)
		return Conn, err
	}

	Conn.SetMaxIdleConns(100)
	Conn.SetMaxOpenConns(100)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err = Conn.PingContext(ctx)
	if err != nil {
		fmt.Printf("Error connecting to the database: %s", err)
		return Conn, err
	}

	fmt.Println("Database connected", "Postgres")
	return Conn, err
}
