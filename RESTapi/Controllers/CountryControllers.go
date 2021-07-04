package Controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ksw95/OMH_BEx/RESTapi/Models"
)

// Create new Country entry.
func (server *Server) CreateCountry(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println("Error reading request body")
			// Error response
		}
		var newCountry Models.Country
		err = json.Unmarshal(reqBody, &newCountry)
		if err != nil {
			fmt.Println("Error unmarshalling request body")
			// Error response
		}
		// Sanitize and Validate inputs
		newCountry.Sanitize()
		err = newCountry.Validate()
		if err != nil {
			fmt.Println("Invalid input")
			// Error response
		}
		result := server.Db.Create(&newCountry)
		if result.Error != nil {
			fmt.Println("Failed to create new country entry in database")
			//error
			return
		}
		// Status created
	}
}

// Show all countries' info stored in the database.
func (server *Server) ShowCountries(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		var countries []Models.Country
		err := server.Db.Find(&countries).Error
		if err != nil {
			fmt.Println("Countries not found")
			// Error response
			return
		}
		// Status 200 with countries info
	}
}

// Show a specific country info.
func (server *Server) ShowCountry(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		params := mux.Vars(req)
		cID, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid input")
			// Error response
			return
		}
		var country Models.Country
		err = server.Db.First(&country, cID).Error
		if err != nil {
			fmt.Println("No country found")
		}
		// status response with country info
	}
}

// Update a specific country info.
func (server *Server) UpdateCountry(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		params := mux.Vars(req)
		cID, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid input")
			// Error response
			return
		}
		new := 0
		var country Models.Country
		err = server.Db.First(&country, cID).Error
		if err != nil {
			fmt.Println("No country found")
			new = 1
		}
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println("Error reading request body")
			// Error response
		}

		var newCountryInfo Models.Country
		err = json.Unmarshal(reqBody, &newCountryInfo)
		if err != nil {
			fmt.Println("Error unmarshalling request body")
		}

		// Sanitize and Validate the inputs
		newCountryInfo.Sanitize()
		err = newCountryInfo.Validate()
		if err != nil {
			fmt.Println("Invalid input")
			// Error response
		}

		if new == 1 {
			result := server.Db.Create(&newCountryInfo)
			if result.Error != nil {
				fmt.Println("Failed to create new country entry in database")
				//error
				return
			}
		} else {
			country.Country = newCountryInfo.Country
			err = server.Db.Save(&country).Error
			if err != nil {
				fmt.Println("Failed to update property info")
				// Error message
				return
			}
		}
		// status ok
	}

	// status not found
}

// Delete country entry from database.
func (server *Server) DeleteCountry(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		params := mux.Vars(req)
		cID, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid input")
		}
		var country Models.Country
		err = server.Db.Find(&country, cID).Error
		if err != nil {
			fmt.Println("Country does not exist")
			// error response
			return
		}
		err = server.Db.Delete(&country, cID).Error
		if err != nil {
			fmt.Println("Failed to delete country entry")
			// error response
			return
		}
		// status 200 ok
	}
}
