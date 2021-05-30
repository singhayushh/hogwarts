package backend

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/singhayushh/hogwarts/backend/config"
	"github.com/singhayushh/hogwarts/backend/services"
)

var server = config.Server{}

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't access .env file")
	}

	services.InitDB(os.Getenv("db_name"), os.Getenv("db_user"), os.Getenv("db_pass"), os.Getenv("db_type"), os.Getenv("db_host"), os.Getenv("db_port"))

	server.Init(os.Getenv("port"))
}
