package controllers

import (
	"github.com/badoux/checkmail"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/maulerrr/book-addict-server/server/DB"
	"github.com/maulerrr/book-addict-server/server/DTO"
	"github.com/maulerrr/book-addict-server/server/helpers"
	"github.com/maulerrr/book-addict-server/server/models"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func Login(context *gin.Context) {
	credentials := new(DTO.LoginDto)
	if err := context.BindJSON(credentials); err != nil {
		helpers.SendMessageWithStatus(context, "Invalid JSON", 400)
		return
	}

	user := models.User{}
	query := models.User{Email: credentials.Email}
	err := DB.DB.First(&user, &query).Error

	if err == gorm.ErrRecordNotFound {
		helpers.SendMessageWithStatus(context, "User not found", 404)
		return
	}

	if !helpers.ComparePasswords(user.Password, credentials.Password) {
		helpers.SendMessageWithStatus(context, "Passwords does not match", 403)
		return
	}

	tokenString, err := generateToken(user)

	if err != nil {
		helpers.SendMessageWithStatus(context, "Auth error (token creation)", 500)
		return
	}

	response := &models.TokenResponse{
		UserID:   user.UserID,
		FullName: user.FullName,
		Email:    user.Email,
		Token:    tokenString,
	}

	helpers.SendSuccessJSON(context, response)
}

func Signup(context *gin.Context) {
	json := new(DTO.SignupDto)
	if err := context.BindJSON(json); err != nil {
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

	err = DB.DB.Create(&newUser).Error
	if err != nil {
		helpers.SendMessageWithStatus(context, err.Error(), 400)
		return
	}

	tokenString, err := generateToken(newUser)

	if err != nil {
		helpers.SendMessageWithStatus(context, "Auth error (token creation)", 500)
		return
	}

	response := &models.TokenResponse{
		UserID:   newUser.UserID,
		FullName: newUser.FullName,
		Email:    newUser.Email,
		Token:    tokenString,
	}

	helpers.SendSuccessJSON(context, response)
}

func generateToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := &models.Claims{
		ID:       user.UserID,
		Email:    user.Email,
		FullName: user.FullName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	jwtKey := os.Getenv("JWT_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	log.Println("jwtkey is: " + (jwtKey))
	log.Println([]byte(jwtKey))

	return tokenString, err
}
