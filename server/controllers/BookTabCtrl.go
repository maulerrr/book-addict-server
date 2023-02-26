package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maulerrr/book-addict-server/server/DB"
	"github.com/maulerrr/book-addict-server/server/DTO"
	"github.com/maulerrr/book-addict-server/server/helpers"
	"github.com/maulerrr/book-addict-server/server/models"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func AddBookTab(context *gin.Context) {
	json := new(DTO.CreateBookTabDto)
	if err := context.ShouldBindJSON(json); err != nil {
		helpers.SendMessageWithStatus(context, "Invalid JSON", 400)
		return
	}

	//add some error handling, validation for book maybe

	newBookTab := models.BookTab{
		BookID:   json.BookID,
		UserID:   json.UserID,
		Finished: false,
	}

	log.Println(newBookTab)

	found := models.BookTab{}
	query := models.BookTab{
		BookID: json.BookID,
		//search field, uniqueness
	}
	err := DB.DB.First(&found, &query).Error
	if err != gorm.ErrRecordNotFound {
		helpers.SendMessageWithStatus(context, "BookTab already exists", 400)
		return
	}

	DB.DB.Create(&newBookTab)

	context.JSON(200, newBookTab)
}

func DeleteBookTabById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		helpers.SendMessageWithStatus(context, "Invalid ID format", 400)
		return
	}

	bookTab, query := models.BookTab{}, models.BookTab{BookID: id}

	err = DB.DB.First(&bookTab, &query).Error
	if err == gorm.ErrRecordNotFound {
		helpers.SendMessageWithStatus(context, "BookTab not found", 400)
		return
	}

	DB.DB.Model(&bookTab).Association("booktabs").Delete()
	DB.DB.Delete(&bookTab)
	context.JSON(200, gin.H{})
}

func GetFavorites(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		helpers.SendMessageWithStatus(context, "Invalid User ID", 400)
		return
	}

	var favorites []models.BookTab
	result := DB.DB.Where("user_id = ? AND finished = false", id).Find(&favorites)
	if result.Error != nil {
		helpers.SendMessageWithStatus(context, result.Error.Error(), 500)
		return
	}
	context.JSON(http.StatusOK, favorites)
}

func GetFinishedBooks(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		helpers.SendMessageWithStatus(context, "Invalid User ID", 400)
		return
	}

	var finishedBooks []models.BookTab
	result := DB.DB.Where("user_id = ? AND finished = true", id).Find(&finishedBooks)
	if result.Error != nil {
		helpers.SendMessageWithStatus(context, result.Error.Error(), 500)
		return
	}
	context.JSON(200, finishedBooks)
}

func GetAllBookTabs(context *gin.Context) {
	bookTabs := []models.BookTab{}
	DB.DB.Preload("book_tabs").Find(&bookTabs)

	context.JSON(200, bookTabs)
}

func UpdateBookTab(context *gin.Context) {
	var requestBody struct {
		UserID   int  `json:"user_id"`
		BookID   int  `json:"book_id"`
		Finished bool `json:"finished"`
	}
	if err := context.BindJSON(&requestBody); err != nil {
		helpers.SendMessageWithStatus(context, "Invalid request body", 400)
		return
	}

	var bookTab models.BookTab
	if err := DB.DB.Where(
		"user_id = ? AND book_id = ?",
		requestBody.UserID,
		requestBody.BookID,
	).First(&bookTab).Error; err != nil {
		helpers.SendMessageWithStatus(context, "BookTab not found", 404)
		return
	}

	bookTab.Finished = requestBody.Finished
	if err := DB.DB.Save(&bookTab).Error; err != nil {
		helpers.SendMessageWithStatus(context, "Failed to update BookTab", 500)
		return
	}

	context.JSON(200, bookTab)
}
