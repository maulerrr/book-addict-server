package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/maulerrr/book-addict-server/server/DB"
	"github.com/maulerrr/book-addict-server/server/helpers"
	"github.com/maulerrr/book-addict-server/server/routes"
)

func main() {
	helpers.SetEnvironment()

	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://127.0.0.1"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	routes.InitRoutes(app)

	DB.ConnectDB()

	log.Fatal(app.Run(port()))
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3001"
	} else {
		port = ":" + port
	}

	return port
}
