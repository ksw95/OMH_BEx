package Database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Builds database source and returns a pointer to the database.
func GetDb() *gorm.DB {
	DbDriver := "mysql"

	// Retrieve values from .env file
	DbHost := os.Getenv("MYSQL_HOSTNAME")
	DbPort := os.Getenv("MYSQL_PORT")
	DbUser := os.Getenv("MYSQL_USER")
	DbPassword := os.Getenv("MYSQL_PASSWORD")
	DbName := os.Getenv("MYSQL_DBNAME")

	// Build Database source string
	DbSource := DbUser + ":" + DbPassword + "@tcp(" + DbHost + ":" + DbPort + ")/" + DbName +
		"?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(DbSource)
	Db, err := gorm.Open(DbDriver, DbSource)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", DbDriver)
		log.Fatal("Error:", err)
	} else {
		fmt.Printf("Successfully connected to the %s database\n", DbDriver)
	}
	return Db
}
