package data

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// "strings"
	"fmt"
	"time"
)

type detailHistory struct {
	VideoId   string `json:"video_id"`
	WatchDate string `json:"watch_date"`
	Title     string `json:"title"`
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Message string `json:"message"`
}

//	@Summary	指定された日付の視聴履歴を返す
//	@Produce	json
//	@Param		date	path		string	true	"yyyy-mm-dd"
//	@Success	200		{array}		detailHistory
//	@Failure	400		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/history/{date} [get]
func GetHistory(c *gin.Context) {
	date := c.Param("date")
	t1, err := time.Parse("2006-01-02", date)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid date format. date format: yyyy-mm-dd",
		})
		return
	}
	fmt.Println(t1)
	t1 = t1.Add(5 * time.Hour)
	t2 := t1.Add(24 * time.Hour)

	rows, err := Db.Query("select A.video_id, A.watch_date, B.title from history as A inner join video_basic_info as B on A.video_id=B.video_id where A.watch_date>=$1 and A.watch_date<=$2 order by A.watch_date desc",
		t1.Format("2006-01-02 15:04:05"), t2.Format("2006-01-02 15:04:05"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// make list of detailHistory
	history := make([]detailHistory, 0)
	for rows.Next() {
		var d detailHistory
		err = rows.Scan(&d.VideoId, &d.WatchDate, &d.Title)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{
				Code: http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
		history = append(history, d)

	}
	rows.Close()
	c.IndentedJSON(http.StatusOK, history)
}

//	@Summary	全ての視聴履歴を返す
//	@Produce	json
//	@Success	200	{array}		detailHistory
//	@Failure	500	{object}	ErrorResponse
//	@Router		/history [get]
func GetAllHistory(c *gin.Context) {
	rows, err := Db.Query("select A.video_id, A.watch_date, B.title from history as A inner join video_basic_info as B on A.video_id=B.video_id order by A.watch_date desc")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// make list of detailHistory
	history := make([]detailHistory, 0)
	for rows.Next() {
		var d detailHistory
		err = rows.Scan(&d.VideoId, &d.WatchDate, &d.Title)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{
				Code: http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
		history = append(history, d)

	}
	rows.Close()
	c.IndentedJSON(http.StatusOK, history)
}
