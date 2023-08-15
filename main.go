package main

import (
	"log"
	"os"
	"test_task/api"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	srv := &api.Server{}

	dsn := os.Getenv("DB_DSN")

	srv.InitDb(dsn)
	srv.InitGin()

	srv.RegisterRoutes()

	srv.Start(":8080")
}
