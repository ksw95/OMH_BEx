package RESTapi

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ksw95/OMH_BEx/RESTapi/Controllers"
	"github.com/ksw95/OMH_BEx/RESTapi/Database"
)

var server = Controllers.Server{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("unable to load values from env file")
	} else {
		log.Println("successfully loaded the values from env file")
	}
}

func Start() {
	server.Initialize()
	Database.Load(server.Db)
	server.RunServer(":8080")
}
