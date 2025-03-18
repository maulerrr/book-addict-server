package models

import "github.com/golang-jwt/jwt/v5"

type TokenResponse struct {
	UserID   int    `json:"userID"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Token    string `json:"token"`
}

type Claims struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`

	jwt.RegisteredClaims
}
