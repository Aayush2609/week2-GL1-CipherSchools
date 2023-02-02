package handler

import (
	"log"
	"net/http"

	"github.com/Aayush2609/week2-GL1-CipherSchools/Golang-application/database"
	"github.com/Aayush2609/week2-GL1-CipherSchools/Golang-application/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


type Handler struct {
	DB *gorm.DB
}




func (h *Handler) GetBooks(in *gin.Context){
	books,err:=database.GetBooks(h.DB)
	if err!=nil{
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, books)

}



func (h *Handler) SaveBook(in *gin.Context){
	book:=models.Book{}
	err := in.BindJSON(&book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.SaveBook(h.DB, &book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, book)

}