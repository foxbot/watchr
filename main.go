package main

import (
	"log"

	"github.com/foxbot/watchr/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	log.Fatal(server.Run())
}
