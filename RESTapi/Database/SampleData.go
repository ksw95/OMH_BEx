package Database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/ksw95/OMH_BEx/RESTapi/Models"
)

var countries = []Models.Country{}

var properties = []Models.Property{}

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
