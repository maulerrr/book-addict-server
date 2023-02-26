package controllers

import (
	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/maulerrr/book-addict-server/server/DB"
	"github.com/maulerrr/book-addict-server/server/DTO"
	"github.com/maulerrr/book-addict-server/server/helpers"
	"github.com/maulerrr/book-addict-server/server/models"
	"gorm.io/gorm"
	"strconv"
)

func GetAllUsers(context *gin.Context) {
	users := []models.User{}
	DB.DB.Preload("users").Find(&users)

	context.JSON(200, users)
}

func GetUserById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		helpers.SendMessageWithStatus(context, "Invalid ID Format", 400)
		return
	}

	user := models.User{}
	query := models.User{UserID: id}
	err = DB.DB.First(&user, &query).Error

	if err == gorm.ErrRecordNotFound {
		helpers.SendMessageWithStatus(context, "User not found", 404)
		return
	}

	context.JSON(200, user)
}

func CreateUser(context *gin.Context) {
	json := new(DTO.CreateUserDto)
	if err := context.ShouldBindJSON(json); err != nil {
		helpers.SendMessageWithStatus(context, "Invalid JSON", 400)
		return
	}

	password := helpers.HashPassword([]byte(json.Password))
	err := checkmail.ValidateFormat(json.Email)
	if err != nil {
		helpers.SendMessageWithStatus(context, "Invalid Email Address", 400)
		return
	}

	newUser := models.User{
		Password: password,
		Email:    json.Email,
		FullName: json.FullName,
	}

	found := models.User{}
	query := models.User{Email: json.Email}
	err = DB.DB.First(&found, &query).Error
	if err != gorm.ErrRecordNotFound {
		helpers.SendMessageWithStatus(context, "User already exists", 400)
		return
	}

	DB.DB.Create(&newUser)

	context.JSON(200, newUser)
}

func DeleteUserById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		helpers.SendMessageWithStatus(context, "Invalid ID format", 400)
		return
	}

	user := models.User{}
	query := models.User{UserID: id}

	err = DB.DB.First(&user, &query).Error
	if err == gorm.ErrRecordNotFound {
		helpers.SendMessageWithStatus(context, "User not found", 400)
		return
	}

	DB.DB.Model(&user).Association("users").Delete()
	DB.DB.Delete(&user)
	context.JSON(200, gin.H{})
}
