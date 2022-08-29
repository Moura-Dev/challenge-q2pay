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

func StartTransaction() (*sqlx.Tx, error) {
	ctx := context.Background()
	tx, err := Conn.BeginTxx(ctx, nil)
	if err != nil {
		fmt.Printf("Error connecting to the database: %s", err)
		return tx, err
	}
	return tx, err
}

func CommitTransaction(tx *sqlx.Tx) error {
	err := tx.Commit()
	if err != nil {
		fmt.Printf("Error connecting to the database: %s", err)
		return err
	}
	return err
}

func RollbackTransaction(tx *sqlx.Tx) error {
	err := tx.Rollback()
	if err != nil {
		fmt.Printf("Error connecting to the database: %s", err)
		return err
	}
	return err
}
