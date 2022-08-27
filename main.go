package main

import (
	"challenge-q2pay/db"
	"challenge-q2pay/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := db.StartDB()

	defer db.Close()
	s := server.NewServer()

	s.Run()
}
