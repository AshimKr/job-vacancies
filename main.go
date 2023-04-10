package main

import (
	"example/challenge/controllers"
	"example/challenge/initializers"
	"example/challenge/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.POST("/job", middleware.RequireAuth, controllers.JobCreate)       //to create a single job in db
	r.POST("/alljob", middleware.RequireAuth, controllers.AllJobCreate) //to create multiple at a time job in db
	r.GET("/job", middleware.RequireAuth, controllers.GetAllJob)        //TO get all the job vacancies
	r.POST("/jobbycountry", middleware.RequireAuth, controllers.JobByCountry)
	r.POST("/jobbycountryorcity", middleware.RequireAuth, controllers.JobByCountryORCity)

	r.Run()
}
