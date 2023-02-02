package database

import (
	"github.com/Aayush2609/week2-GL1-CipherSchools/Golang-application/models"
	"github.com/jinzhu/gorm"
)

//GetBooks is creating connection and interacting from golang app to db server via db variable

func GetBooks(db *gorm.DB) ([]models.Book, error) {

	books := []models.Book{}
	query := db.Select("books.*")
	err :=query.Find(&books).Error  //all the books will be populated
	if err != nil {
		return nil, err
	}
	
	
	return books, nil
}

func GetBookByID(db *gorm.DB, id string) (*models.Book, error){
	book:=models.Book{}
	err := db.Select("books.*").Group("books.id").Where("books.id= ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}
func DeleteBookByID(db *gorm.DB, id string) error {
	var book models.Book
	err := db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateBookByID(db *gorm.DB, book *models.Book) error {
	err := db.Save(book).Error //book is with id
	if err != nil {
		return err
	}
	return nil
}
func SaveBook(db *gorm.DB, book *models.Book) error {
	err := db.Save(book).Error //book here is without id
	if err != nil {
		return err
	}
	return nil
}