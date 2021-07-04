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
		err := server.Db.Where("Available = ?", "Yes").Find(&properties).Error
		if err != nil {
			fmt.Println("Properties not found")
			// Error response
			return
		}
		// status response 200 with avail properties info
	}
}

// Show properties based on specific countries.
// Takes in country ID and returns info of all properties in said country.
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
			// Error response
			return
		}
		var properties []Models.Property
		err = server.Db.Where("Country = ?", country.Country).Find(&properties).Error
		if err != nil {
			fmt.Println("Properties not found")
			// Error response
			return
		}
		// Status response 200 with properties info
	}
}

// Show a chosen property information.
// Takes in a property ID and return property info with said ID.
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

// Update an existing property entry.
// Allows changes to Address, Country, Description and Availiability.
func (server *Server) UpdateProperty(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		params := mux.Vars(req)
		pID, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid input as property ID")
			// Error response
			return
		}

		new := 0
		var oldPropInfo Models.Property
		err = server.Db.First(&oldPropInfo, pID).Error
		if err != nil {
			fmt.Println("Property does not exist")
			new = 1
		}

		var newPropInfo Models.Property
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println("Error reading request body")
			// Error response
			return
		}

		err = json.Unmarshal(reqBody, &newPropInfo)
		if err != nil {
			fmt.Println("Failed to unmarshal request body")
			// Error reponse
			return
		}

		// Sanitize updated info
		newPropInfo.Sanitize()

		// Check for changes in country and whether country exists
		var country Models.Country
		err = server.Db.Where("Country = ?", newPropInfo.Country).Find(&country).Error
		if err != nil {
			fmt.Println("Country not found")
			// Create new Country entry
		}

		if new == 1 {
			// Validate new property inputs
			err = newPropInfo.Validate()
			if err != nil {
				fmt.Println("Invalid input")
				//error
				return
			}
			result := server.Db.Create(&newPropInfo)
			if result.Error != nil {
				fmt.Println("Failed to create new property entry in database")
				//error
				return
			}
			// statusCreated
		} else {
			if newPropInfo.Address != "" {
				oldPropInfo.Address = newPropInfo.Address
			}
			if newPropInfo.Country != "" {
				oldPropInfo.Country = newPropInfo.Country
			}
			if newPropInfo.Description != "" {
				oldPropInfo.Description = newPropInfo.Description
			}
			if newPropInfo.Available != "" {
				oldPropInfo.Available = newPropInfo.Available
			}
			err = oldPropInfo.Validate()
			if err != nil {
				fmt.Println("Invalid input")
				// Error response
				return
			}
			err = server.Db.Save(&oldPropInfo).Error
			if err != nil {
				fmt.Println("Failed to update property info")
				// Error message
				return
			}
		}
		// Status 200 okay message
	}
}

// Delete a property entry inside database.
func (server *Server) DeleteProperty(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		params := mux.Vars(req)
		pID, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid input as property ID")
			// Error response
			return
		}
		var propInfo Models.Property
		err = server.Db.First(&propInfo, pID).Error
		if err != nil {
			fmt.Println("Property does not exist")
			// Error response
			return
		}
		err = server.Db.Delete(&propInfo, pID).Error
		if err != nil {
			fmt.Println("Failed to delete property entry")
			// Error Response
		}

		// Success response
	}
}
