package main

import (
	"log"
	"wan-api-kol-event/Controllers"
	"wan-api-kol-event/Initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	Initializers.LoadEnvironmentVariables()
	Initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Define your Gin routes here
	r.GET("/kols", Controllers.GetKolsController)
	
	// Define the route for generating dummy data using POST method
	r.POST("/generate-dummy-data", Controllers.GenerateDummyData)

	// Run Gin server
	if err := r.Run(":8081"); err != nil {
		log.Println("Failed to start server")
	}
}
