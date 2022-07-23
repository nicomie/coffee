package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "vue-postgres.json")
	err := godotenv.Load("envs/app.env", "envs/postgres.env")
	if err != nil {
		log.Fatalf("Error loading conf: %s", err)
	}
}
