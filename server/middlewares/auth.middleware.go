package middlewares

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/maulerrr/book-addict-server/server/helpers"
	"github.com/maulerrr/book-addict-server/server/models"
)

func parseToken(tokenString string) (*models.Claims, error) {
	jwtKey := os.Getenv("JWT_KEY")
	claims := &models.Claims{}

	log.Println(jwtKey)

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			helpers.SendMessageWithStatus(context, "Authorize!", 404)
			context.Abort()
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			helpers.SendMessageWithStatus(context, "Incorrect authorization header", 403)
			context.Abort()
			return
		}

		token := headerParts[1]

		jwtKey := os.Getenv("JWT_KEY")
		claims := &models.Claims{}

		tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				helpers.SendMessageWithStatus(context, "Unauthorized", 401)
				context.Abort()
				return
			}

			helpers.SendMessageWithStatus(context, err.Error(), 404)
			context.Abort()
			return
		}

		if !tkn.Valid {
			helpers.SendMessageWithStatus(context, "Unauthorized", 401)
			context.Abort()
			return
		}

		context.Next()
	}
}
