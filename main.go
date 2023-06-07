package main

import (
	postgres "ass_3/pkg/store/postgres"
	"ass_3/services/contact/internal/delivery"
	"ass_3/services/contact/internal/repository"
	"ass_3/services/contact/internal/useCase"
	"net/http"

	// "ass_3/services/contact/internal/repository"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	host, port, user, password, dbname := envInit()
	db, err := postgres.OpenDb(host, sport, user, password, dbname)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Print("Connection to db succesfully established!")

	defer db.Close()
	newRepository := repository.NewRepository()
	newUseCase := useCase.NewUseCase()
	newDelivery := delivery.NewDelivery(newUseCase, newRepository)

	fmt.Println("Repository: ", newRepository)
	fmt.Println("UseCase: ", newUseCase)
	fmt.Println("Delivery: ", newDelivery)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func envInit() (string, int, string, string, string) {

	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	return host, port, user, password, dbname
}
