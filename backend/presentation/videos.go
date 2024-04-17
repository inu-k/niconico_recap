package presentation

import (
	"net/http"

	"niconico_recap_backend/logic"

	"github.com/gin-gonic/gin"
)

// @Summary	指定されたvideoIdの動画情報を返す
// @Produce	json
// @Param		videoId	path		string	true	"videoId"
// @Success	200		{object}	data.VideoInfo
// @Failure	404		{object}	ErrorResponse
// @Failure	500		{object}	ErrorResponse
// @Router		/videos/{videoId} [get]
func GetVideoInfo(c *gin.Context) {
	videoId := c.Param("videoId")
	videoInfo, err := logic.FetchVideoInfo(videoId)
	if err != nil {
		statusCode := http.StatusInternalServerError

		if err.Error() == "video not found" {
			statusCode = http.StatusNotFound
		}

		c.IndentedJSON(statusCode, ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, videoInfo)
}
