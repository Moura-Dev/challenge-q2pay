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
	database, err := db.StartDB()
	database.MustExec(db.Schema)

	defer database.Close()
	s := server.NewServer()

	s.Run()
}
