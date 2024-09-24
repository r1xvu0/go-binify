package models

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomClaims struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type Paste struct {
	gorm.Model
	Title    string
	Content  string
	AuthorID uint
}
