package RESTapi

var server = Controllers.Server{}

func Start() {
	server.Initialize()
	Sample.Load(server.Db)
	server.RunServer(":8080")
}
