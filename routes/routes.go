package routes

import (
	"github.com/ananascharles/binify/database"
	"github.com/ananascharles/binify/handlers"
	"github.com/ananascharles/binify/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	db, err := database.SetupDB()
	if err != nil {
		panic("DB is not setup correctly")
	}

	database.MigrateDB(db)

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"http://gotent"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// router.SetTrustedProxies([]string{"http://gotent"})

	router.GET("/", handlers.HandleIndex)
	router.GET("/login", handlers.LoginHandler)
	router.GET("/protected", middleware.AuthMiddleware(), handlers.ProtectedHandler)
	router.POST("/createPaste", middleware.AuthMiddleware(), func(c *gin.Context) {
		handlers.CreatePasteHandler(c, db)
	})
	router.GET("/getPastes", middleware.AuthMiddleware(), func(c *gin.Context) {
		handlers.GetAllPastesHandler(c, db)
	})
	router.GET("/paste", func(c *gin.Context) {
		handlers.GetPasteHandler(c, db)
	})

	return router
}
