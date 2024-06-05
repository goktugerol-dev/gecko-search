package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env := godotenv.Load()

	if env != nil {
		panic("Cannot find environment variables")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":4000"
	} else {
		port = ":" + port
	}
}
