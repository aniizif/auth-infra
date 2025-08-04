package main

import (
	"github.com/aniizif/stack-mate/auth-service/internal/app"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	_ = godotenv.Load()
	application := app.NewApp()
	log.Println("Starting Auth Service...")
	application.Run()
}
