package main

import (
	"log"

	"github.com/Aayush2609/week2-GL1-CipherSchools/Golang-application/database"
	"github.com/Aayush2609/week2-GL1-CipherSchools/Golang-application/routers"
)
func main() {
	database.Setup()  //establish the database connection
	engine:=routers.Router()  //get the customized engine
	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}



