package Controllers

func (server *Server) InitializeRoutes() {
	URI := "/api/v1"

	// Country Routes
	server.Router.HandleFunc(URI+"", server.CreateCountry).Methods("POST")
	server.Router.HandleFunc(URI+"", server.ShowCountries).Methods("GET")
	server.Router.HandleFunc(URI+"", server.ShowCountry).Methods("GET")
	server.Router.HandleFunc(URI+"", server.UpdateCountry).Methods("PUT")
	server.Router.HandleFunc(URI+"", server.DeleteCountry).Methods("DELETE")

	// Property Routes
	server.Router.HandleFunc(URI+"", server.CreateProperty).Methods("POST")
	server.Router.HandleFunc(URI+"", server.AllProperties).Methods("GET")
	server.Router.HandleFunc(URI+"", server.AvailProperties).Methods("GET")
	server.Router.HandleFunc(URI+"", server.CountryProperties).Methods("GET")
	server.Router.HandleFunc(URI+"", server.ViewProperty).Methods("GET")
	server.Router.HandleFunc(URI+"", server.UpdateProperty).Methods("PUT")
	server.Router.HandleFunc(URI+"", server.DeleteProperty).Methods("DELETE")
}
