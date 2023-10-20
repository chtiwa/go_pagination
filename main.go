package main

import (
	"github.com/chtiwa/go_pagination/controllers"
	"github.com/chtiwa/go_pagination/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDB()
}

func main() {
	app := gin.Default()

	app.LoadHTMLGlob("templates/**/*")
	app.Static("/assets", "./assets")

	app.GET("/", controllers.PeopleIndexGET)
	app.GET("/page/:page", controllers.PeopleIndexGET)

	app.Run()
}
