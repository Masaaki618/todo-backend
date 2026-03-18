package infra

import "github.com/joho/godotenv"

func Initialize() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
