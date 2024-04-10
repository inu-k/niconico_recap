package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"niconico_recap_backend/data"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
	docs "niconico_recap_backend/docs"
)

func pong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

//	@title		niconico_recap_backend API
//	@version	0.1
//	@license	inu-k
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/ping", pong)
	r.GET("/history", data.GetAllHistory)
	r.GET("/history/:date", data.GetHistory)
	r.GET("/videos/:videoId", data.GetVideo)
	r.GET("/summary/daily/:date", data.GetDailySummary)
	r.GET("/summary", data.GetSummary)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080")
}