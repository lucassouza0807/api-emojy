package main

import (
	logincontroller "api-emoji/internal/controllers/login"
	phrasecontroller "api-emoji/internal/controllers/phraseController"
	userController "api-emoji/internal/controllers/userController"
	authMiddleware "api-emoji/internal/middlewares/authMiddleware"
	"api-emoji/internal/utils/migration"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initLogger() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("erro ao criar arquivo de log: %v", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	initLogger()
	//Sinc all tables
	migration.RunMigration()

	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:8080",
			"http://192.168.0.124:3000",
			"http://192.168.0.124:8080",
		},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	app.GET("/api/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Tudo ok",
		})
	})
	//Login
	app.POST("/api/v1/create-user", userController.CreateUser)
	app.POST("/api/v1/login", logincontroller.Login)

	//Protected Routes
	api := app.Group("/api/v1/", authMiddleware.VerifyToken)
	{
		//PhrasesControllers handlers
		api.GET("/phrases", phrasecontroller.GetUserPhrases)
		api.GET("/search-phrase", phrasecontroller.SearchForPhrases)
		api.POST("store-phrase", phrasecontroller.CreatePhrase)
		api.PUT("/edit-phrase/:id", phrasecontroller.EditPhrase)
		api.DELETE("/delete-phrase/:id", phrasecontroller.DeletePhrase)
	}

	app.Run()
}
