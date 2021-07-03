package Database

import (
	"log"

	models "github.com/ksw95/OMH_BEx/RESTapi/Models"

	"github.com/jinzhu/gorm"
)

var countries = []models.Country{}

var properties = []models.Property{}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Country{}, &models.Property{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v\n", err)
	}

	err = db.Debug().AutoMigrate(&models.Country{}, &models.Property{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v\n", err)
	}

	for i := range countries {
		err = db.Debug().Model(&models.Country{}).Create(&countries[i]).Error
		if err != nil {
			log.Fatalf("cannot seed coments table %v\n", err)
		}
	}

	for i := range properties {
		err = db.Debug().Model(&models.Property{}).Create(&properties[i]).Error
		if err != nil {
			log.Fatalf("cannot seed coments table %v\n", err)
		}
	}
}
