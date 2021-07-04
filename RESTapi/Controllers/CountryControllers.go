package Controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ksw95/OMH_BEx/RESTapi/Models"
	"github.com/ksw95/OMH_BEx/RESTapi/Responses"
)

// Create new Country entry.
func (server *Server) CreateCountry(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println("Error reading request body")
			Responses.ErrRes(res, http.StatusInternalServerError, err)
			return
		}
		var newCountry Models.Country
		err = json.Unmarshal(reqBody, &newCountry)
		if err != nil {
			fmt.Println("Error unmarshalling request body")
			Responses.ErrRes(res, http.StatusInternalServerError, err)
			return
		}
		// Sanitize and Validate inputs
		newCountry.Sanitize()
		err = newCountry.Validate()
		if err != nil {
			fmt.Println("Invalid input")
			Responses.ErrRes(res, http.StatusInternalServerError, err)
			return
		}
		result := server.Db.Create(&newCountry)
		if result.Error != nil {
			fmt.Println("Failed to create new country entry in database")
			Responses.ErrRes(res, http.StatusInternalServerError, err)
			return
		}
		Responses.BasicRes(res, http.StatusCreated, newCountry)
	}
}

// Show all countries' info stored in the database.
func (server *Server) ShowCountries(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		var countries []Models.Country
		err := server.Db.Find(&countries).Error
		if err != nil {
			fmt.Println("Countries not found")
			Responses.ErrRes(res, http.StatusInternalServerError, err)
			return
		}
		Responses.BasicRes(res, http.StatusOK, countries)
	}
}

// Show a specific country info.
func (server *Server) ShowCountry(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		params := mux.Vars(req)
		cID, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid input")
			Responses.ErrRes(res, http.StatusInternalServerError, errors.New("invalid input"))
			return
		}
		var country Models.Country
		err = server.Db.First(&country, cID).Error
		if err != nil {
			fmt.Println("No country found")
			Responses.ErrRes(res, http.StatusInternalServerError, err)
			return
		}
		Responses.BasicRes(res, http.StatusOK, country)
	}
}

// Update a specific country info.
func (server *Server) UpdateCountry(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		params := mux.Vars(req)
		cID, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid input")
			Responses.ErrRes(res, http.StatusInternalServerError, errors.New("invalid input"))
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
			Responses.ErrRes(res, http.StatusInternalServerError, err)
			return
		}

		var newCountryInfo Models.Country
		err = json.Unmarshal(reqBody, &newCountryInfo)
		if err != nil {
			fmt.Println("Error unmarshalling request body")
			Responses.ErrRes(res, http.StatusInternalServerError, err)
			return
		}

		// Sanitize and Validate the inputs
		newCountryInfo.Sanitize()
		err = newCountryInfo.Validate()
		if err != nil {
			fmt.Println("Invalid input")
			Responses.ErrRes(res, http.StatusBadRequest, err)
			return
		}

		if new == 1 {
			result := server.Db.Create(&newCountryInfo)
			if result.Error != nil {
				fmt.Println("Failed to create new country entry in database")
				Responses.ErrRes(res, http.StatusInternalServerError, err)
				return
			}
			Responses.BasicRes(res, http.StatusOK, newCountryInfo)
		} else {
			country.Country = newCountryInfo.Country
			err = server.Db.Save(&country).Error
			if err != nil {
				fmt.Println("Failed to update property info")
				Responses.ErrRes(res, http.StatusInternalServerError, err)
				return
			}
			Responses.BasicRes(res, http.StatusOK, country)
		}
	}
}

// Delete country entry from database.
func (server *Server) DeleteCountry(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-type") == "application/json" {
		params := mux.Vars(req)
		cID, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid input")
			Responses.ErrRes(res, http.StatusInternalServerError, errors.New("invalid input"))
			return
		}
		var country Models.Country
		err = server.Db.Find(&country, cID).Error
		if err != nil {
			fmt.Println("Country does not exist")
			Responses.ErrRes(res, http.StatusInternalServerError, err)
			return
		}
		err = server.Db.Delete(&country, cID).Error
		if err != nil {
			fmt.Println("Failed to delete country entry")
			Responses.ErrRes(res, http.StatusInternalServerError, err)
			return
		}
		Responses.BasicRes(res, http.StatusOK, "Deleted successfully")
	}
}
