package main

import (
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.Use(static.Serve("/", static.LocalFile("./build", true)))
	websiteRouter := app.Group("/")
	websiteRouter.GET("/staking", websiteHandler)
	websiteRouter.GET("/stake", websiteHandler)
	PORT := os.Getenv("PORT")
	app.Run(":" + PORT)
}

func websiteHandler(c *gin.Context) {
	// let react handle the website
	c.File("./build/index.html")
}
