package Controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ksw95/OMH_BEx/RESTapi/Database"
)

type Server struct {
	Db     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {

	server.Db = Database.GetDb()

	server.Router = mux.NewRouter().StrictSlash(true)

	// Prepare routes for multiplexer
	server.InitializeRoutes()
}

func (server *Server) RunServer(addr string) {
	log.Printf("Server is listening on port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
