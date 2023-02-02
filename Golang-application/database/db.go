package database

import (
	"log"
	"github.com/Aayush2609/week2-GL1-CipherSchools/Golang-application/models"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
func GetDB()*gorm.DB {
	return DB
}

func Setup() {
	host :="localhost"
	port :="5433"
	dbName :="book"
	username :="postgres"
	password := "postgres"
	args := "host="+host+" port="+port+" user="+username+" dbname="+dbName+" sslmode=disable password="+password

	db,err:=gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(models.Book{})
	DB = db //giving the control to the db variable
}