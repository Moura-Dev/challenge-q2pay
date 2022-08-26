package main

import (
	"challenge-q2pay/db"
	"challenge-q2pay/server"
)

func main() {
	db.StartDB()

	s := server.NewServer()

	s.Run()
}
