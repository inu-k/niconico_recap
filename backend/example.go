package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"niconico_recap_backend/data"
)

func pong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", pong)
	r.GET("/history", data.GetAllHistory)
	r.GET("/history/:date", data.GetHistory)
	r.GET("/videos/:videoId", data.GetVideo)
	r.GET("/summary/:date", data.GetSummary)

	r.Run(":8080")
}