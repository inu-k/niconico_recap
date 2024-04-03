package data

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoInfo struct {
	VideoId string `json:"video_id"`
	Title   string `json:"title"`
	Tags    []string `json:"tags"`
	ThumbnailUrl string `json:"thumbnail_url"`
}

func GetVideo(c *gin.Context) {
	videoId := c.Param("videoId")
	rows, err := Db.Query("select video_id, title, thumbnail_url from video_basic_info where video_id=$1", videoId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var v VideoInfo
	if rows.Next() {
		err = rows.Scan(&v.VideoId, &v.Title, &v.ThumbnailUrl)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Not found",
		})
		return
	}
	rows.Close()

	tags := make([]string, 0)

	rows, err = Db.Query("select tag from video_tag_info where video_id=$1", videoId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		tags = append(tags, tag)
	}

	v.Tags = tags
	c.IndentedJSON(http.StatusOK, v)
}