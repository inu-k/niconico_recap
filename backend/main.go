package main

import (
	"net/http"
	"niconico_recap_backend/data"
	docs "niconico_recap_backend/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func pong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// @title		niconico_recap_backend API
// @version	0.1
// @license	TBD
func main() {
	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3030", "http://localhost:3000"},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/ping", pong)
	r.GET("/history", data.GetAllHistory)
	r.GET("/history/:date", data.GetHistory)
	r.GET("/videos/:videoId", data.GetVideo)
	r.GET("/summary", data.GetSummary)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080")
}
