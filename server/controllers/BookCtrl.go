package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maulerrr/book-addict-server/server/DB"
	"github.com/maulerrr/book-addict-server/server/DTO"
	"github.com/maulerrr/book-addict-server/server/helpers"
	"github.com/maulerrr/book-addict-server/server/models"
	"gorm.io/gorm"
	"strconv"
)

func GetAllBooks(context *gin.Context) {
	books := []models.Book{}
	DB.DB.Preload("books").Find(&books)

	context.JSON(200, books)
}

func GetBookById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		helpers.SendMessageWithStatus(context, "Invalid ID Format", 400)
		return
	}

	book, query := models.Book{}, models.Book{ID: id}
	err = DB.DB.First(&book, &query).Error

	if err == gorm.ErrRecordNotFound {
		helpers.SendMessageWithStatus(context, "Book not found", 404)
		return
	}

	context.JSON(200, book)
}

func AddBook(context *gin.Context) {
	json := new(DTO.CreateBookDto)
	if err := context.ShouldBindJSON(json); err != nil {
		helpers.SendMessageWithStatus(context, "Invalid JSON", 400)
		return
	}

	//add some error handling, validation for book maybe

	newBook := models.Book{
		Title:       json.Title,
		Author:      json.Author,
		Price:       json.Price,
		Description: json.Description,
		Img:         json.Img,
	}

	found := models.Book{}
	query := models.Book{
		Title: json.Title,
	}
	err := DB.DB.First(&found, &query).Error
	if err != gorm.ErrRecordNotFound {
		helpers.SendMessageWithStatus(context, "Book already exists", 400)
		return
	}

	DB.DB.Create(&newBook)

	context.JSON(200, newBook)
}

func DeleteBookById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		helpers.SendMessageWithStatus(context, "Invalid ID format", 400)
		return
	}

	book, query := models.Book{}, models.Book{ID: id}

	err = DB.DB.First(&book, &query).Error
	if err == gorm.ErrRecordNotFound {
		helpers.SendMessageWithStatus(context, "Book not found", 400)
		return
	}

	DB.DB.Model(&book).Association("books").Delete()
	DB.DB.Delete(&book)
	context.JSON(200, gin.H{})
}
