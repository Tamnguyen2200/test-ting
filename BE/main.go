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

// CORS middleware
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Allow all sources
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle OPTIONS method
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204) // No Content
            return
        }

        c.Next()
    }
}


func main() {
	r := gin.Default()

	// Apply the CORS middleware
	r.Use(CORSMiddleware())

	// Define your Gin routes here
	r.GET("/kols", Controllers.GetKolsController)

	// Define the route for generating dummy data using POST method
	r.POST("/generate-dummy-data", Controllers.GenerateDummyData)


	// Run Gin server
	if err := r.Run(":8081"); err != nil {
		log.Println("Failed to start server")
	}
}
