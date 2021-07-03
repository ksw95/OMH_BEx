package RESTapi

import (
	"github.com/ksw95/OMH_BEx/RESTapi/Controllers"
	"github.com/ksw95/OMH_BEx/RESTapi/Database"
)

var server = controllers.Server{}

func Start() {
	server.Initialize()
	Database.Load(server.Db)
	server.RunServer(":8080")
}
