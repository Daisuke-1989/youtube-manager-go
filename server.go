package main

import (
	"youtube-manager-go/routes"

	"youtube-manager-go/middlewares"

	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"

	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env")
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middlewares.YouTubeService())
	e.Use(middlewares.DatabaseService())
	e.Use(middlewares.Firebase())

	//Routes
	routes.Init(e)

	//start servers
	e.Logger.Fatal(e.Start(":8080"))
}
