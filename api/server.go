package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kjuchniewicz/fullstack/api/controllers"
	"github.com/kjuchniewicz/fullstack/api/seed"
)

var server = controllers.Server{}

func Run() {
	var err error = godotenv.Load()
	if err != nil {
		log.Fatalf("Błąd podczas pobierania .env, brak przejścia %v", err)
	} else {
		fmt.Println("Otrzymujemy wartości .env")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":8088")
}
