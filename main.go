package main

import (
	"log"
	"os"
	"todo-backend/infra"
	"todo-backend/routers"
)

func main() {
	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = "0.0.0.0:8080"
	}

	infra.Initialize()
	r := routers.NewRouter()
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
