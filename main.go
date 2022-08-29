package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mahfuzan/localswiftoms-v2/controllers"
	"github.com/mahfuzan/localswiftoms-v2/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.GET("/company", controllers.GetCompanyList)
	r.GET("/company/:id", controllers.GetCompanyById)
	r.POST("/company", controllers.CreateCompany)
	r.PUT("/company", controllers.UpdateCompanyById)
	r.DELETE("/company/:id", controllers.DeleteCompanyById)

	r.GET("/location", controllers.GetLocationList)
	r.GET("/location/:id", controllers.GetLocationById)
	// r.POST("/location", controllers.CreateLocation)
	// r.PUT("/location", controllers.UpdateLocationById)
	// r.DELETE("/location/:id", controllers.DeleteLocationById)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
