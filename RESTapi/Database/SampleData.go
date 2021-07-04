package Database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/ksw95/OMH_BEx/RESTapi/Models"
)

var countries = []Models.Country{
	{
		Country: "Singapore",
	},
	{
		Country: "Malaysia",
	},
	{
		Country: "Philipines",
	},
}

var properties = []Models.Property{
	{
		Address:     "535 Clementi Rd, Singapore 599489",
		Country:     "Singapore",
		Description: "Ngee Ann Polytechnic",
		Available:   "No",
	},
	{
		Address:     "3151, Commonwealth Avenue, Singapore 129581",
		Country:     "Singapore",
		Description: "Commercial unit for rent",
		Available:   "Yes",
	},
	{
		Address:     "Sample Test Data 3",
		Country:     "Malaysia",
		Description: "For testing Country Properties display",
		Available:   "Yes",
	},
	{
		Address:     "Sample Test Data 4",
		Country:     "Philipines",
		Description: "For testing Country properties display,",
		Available:   "No",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&Models.Country{}, &Models.Property{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v\n", err)
	}

	err = db.Debug().AutoMigrate(&Models.Country{}, &Models.Property{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v\n", err)
	}

	for i := range countries {
		err = db.Debug().Model(&Models.Country{}).Create(&countries[i]).Error
		if err != nil {
			log.Fatalf("cannot seed coments table %v\n", err)
		}
	}

	for i := range properties {
		err = db.Debug().Model(&Models.Property{}).Create(&properties[i]).Error
		if err != nil {
			log.Fatalf("cannot seed coments table %v\n", err)
		}
	}
}
