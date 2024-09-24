package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ananascharles/binify/globals"
	"github.com/ananascharles/binify/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func LoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "admin" && password == "secret" {
		claims := models.CustomClaims{
			UserID:   1,
			Username: username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(globals.SecretKey))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func ProtectedHandler(c *gin.Context) {
	user := c.MustGet("user").(*models.CustomClaims)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello %s!", user.Username)})
}

func CreatePasteHandler(c *gin.Context, db *gorm.DB) {
	var request struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paste := models.Paste{
		Title:   request.Title,
		Content: request.Content,
	}

	result := db.Create(&paste)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, paste)
}

func GetAllPastesHandler(c *gin.Context, db *gorm.DB) {
	var pastes []models.Paste
	result := db.Find(&pastes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, pastes)
}

func GetPasteHandler(c *gin.Context, db *gorm.DB) {
	var paste models.Paste
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No ID passed"})
		return
	}
	result := db.First(&paste, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%v: ID %s", result.Error.Error(), id)})
		return
	}

	c.JSON(http.StatusOK, paste)
}

func HandleIndex(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"api_name":    "Binify",
		"version":     "0.0.1",
		"description": "Simple yet powerful clone of Pastebin using Go as a backend",
	})
}
