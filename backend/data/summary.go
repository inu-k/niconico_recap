package data

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Summary struct {
	TagCount map[string]int `json:"tag_count"`
}

func GetSummary(c *gin.Context) {
	date := c.Param("date")
	t1, err := time.Parse("2006-01-02", date)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date format\ndate format: yyyy-mm-dd",
		})
		return
	}
	t1 = t1.Add(5 * time.Hour)
	t2 := t1.Add(24 * time.Hour)

	rows, err := Db.Query("select B.tag, count(*) from history as A inner join video_tag_info as B on A.video_id=B.video_id where A.watch_date>=$1 and A.watch_date<=$2 group by B.tag",
		t1.Format("2006-01-02 15:04:05"), t2.Format("2006-01-02 15:04:05"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	summary := Summary{TagCount: make(map[string]int)}
	for rows.Next() {
		var tag string
		var count int
		err = rows.Scan(&tag, &count)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		summary.TagCount[tag] = count
	}
	rows.Close()
	c.IndentedJSON(http.StatusOK, summary)
}