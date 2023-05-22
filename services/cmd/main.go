package main

import (
	postgres "ass_3/pkg/store/postgres"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	host, port, user, password, dbname := envInit()
	db, err := postgres.OpenDb(host, port, user, password, dbname)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Print("Connection to db succesfully established!")

	defer db.Close()

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
