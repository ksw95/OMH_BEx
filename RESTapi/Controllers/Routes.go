package Controllers

func (server *Server) InitializeRoutes() {
	URI := "/api/v1"

	// Country Routes
	server.Router.HandleFunc(URI+"/country", server.CreateCountry).Methods("POST")
	server.Router.HandleFunc(URI+"/country", server.ShowCountries).Methods("GET")
	server.Router.HandleFunc(URI+"/country/{id:[0-9]+}", server.ShowCountry).Methods("GET")
	server.Router.HandleFunc(URI+"/country/{id:[0-9]+}", server.UpdateCountry).Methods("PUT")
	server.Router.HandleFunc(URI+"/country/{id:[0-9]+}", server.DeleteCountry).Methods("DELETE")

	// Property Routes
	server.Router.HandleFunc(URI+"/property", server.CreateProperty).Methods("POST")
	server.Router.HandleFunc(URI+"/property", server.AllProperties).Methods("GET")
	server.Router.HandleFunc(URI+"/property/avail", server.AvailProperties).Methods("GET")
	server.Router.HandleFunc(URI+"/property/country/{id:[0-9]+}", server.CountryProperties).Methods("GET")
	server.Router.HandleFunc(URI+"/property/{id:[0-9]+}", server.ViewProperty).Methods("GET")
	server.Router.HandleFunc(URI+"/property/{id:[0-9]+}", server.UpdateProperty).Methods("PUT")
	server.Router.HandleFunc(URI+"/property/{id:[0-9]+}", server.DeleteProperty).Methods("DELETE")
}
