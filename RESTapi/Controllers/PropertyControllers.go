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

// Create new property entry and store into database.
func (server *Server) CreateProperty(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println("Failed to read request body")
			// error
		} else {
			newProperty := Models.Property{}
			err := json.Unmarshal(reqBody, &newProperty)
			if err != nil {
				fmt.Println("Failed to unmarshal request body")
				//error
				return
			}
			// Sanitize and Validate
			newProperty.Sanitize()
			err = newProperty.Validate()
			if err != nil {
				fmt.Println("Invalid input")
				//error
				return
			}
			result := server.Db.Create(&newProperty)
			if result.Error != nil {
				fmt.Println("Failed to create new property entry in database")
				//error
				return
			}
			// statusCreated
			return
		}
	}
}

// Show all properties in database.
func (server *Server) AllProperties(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		var properties []Models.Property
		err := server.Db.Find(&properties).Error
		if err != nil {
			fmt.Println("Properties not found")
			return
		}
		// status response 200 with properties info
	}
}

// Show all available properties in database.
func (server *Server) AvailProperties(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		var properties []Models.Property
		err := server.Db.Where("Available = ?", true).Find(&properties).Error
		if err != nil {
			fmt.Println("Properties not found")
			return
		}
		// status response 200 with avail properties info
	}
}

// Show properties based on specific countries.
func (server *Server) CountryProperties(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		params := mux.Vars(req)
		cID, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid input as country ID")
			// Error response
			return
		}
		var country Models.Country
		err = server.Db.Where("ID = ?", cID).Find(&country).Error
		if err != nil {
			fmt.Println("Country not found")
			return
		}
		var properties []Models.Property
		err = server.Db.Where("Country = ?", country.Country).Find(&properties).Error
		if err != nil {
			fmt.Println("Properties not found")
			return
		}
		// Status response 200 with properties info
	}
}

// Show a chosen property information.
func (server *Server) ViewProperty(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		params := mux.Vars(req)
		pID, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid input as property ID")
			// Error response
			return
		}
		var property Models.Property
		err = server.Db.First(&property, pID).Error
		if err != nil {
			fmt.Println("Property not found")
		}
		// status response 200 with property info
	}
}
