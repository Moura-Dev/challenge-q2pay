package main

import (
	"challenge-q2pay/db"
	"challenge-q2pay/server"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if os.Getenv("ENV") == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}

	database, _ := db.StartDB()
	database.MustExec(db.Schema)

	defer database.Close()
	s := server.NewServer()

	s.Run()
}
