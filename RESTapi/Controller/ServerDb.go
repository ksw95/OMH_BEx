package Controller

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	Db     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {
	var err error

	DbDriver := "mysql"

	// Retrieve values from .env file
	DbHost := os.Getenv("MYSQL_HOSTNAME")
	DbPort := os.Getenv("MYSQL_PORT")
	DbUser := os.Getenv("MYSQL_USER")
	DbPassword := os.Getenv("MYSQL_PASSWORD")
	DbName := os.Getenv("MYSQL_DBNAME)")

	// Build Database source string
	DbSource := DbUser + ":" + DbPassword + "@tcp(" + DbHost + ":" + DbPort + ")/" + DbName +
		"?charset=utf8&parseTime=True&loc=Local"

	server.Db, err = gorm.Open(DbDriver, DbSource)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", DbDriver)
		log.Fatal("Error:", err)
	} else {
		fmt.Printf("Successfully connected to the %s database\n", DbDriver)
	}

	server.Router = mux.NewRouter().StrictSlash(true)

	// Prepare routes for multiplexer
	server.InitializeRoutes()
}

func (server *Server) RunServer(addr string) {
	log.Printf("Server is listening on port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
