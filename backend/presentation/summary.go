package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"niconico_recap_backend/logic"
	"time"
)

// @Description	指定された期間内に視聴された動画のタグのサマリーを返す
// @Description	タグは視聴された回数の降順で返される
// @Description	dateが指定された場合はその日の5時から29時までのデータを返す
// @Description	startDateとendDateが指定された場合はstartDateの5時からからendDateの29時までのデータを返す
// @Description	startDateが指定されない場合は1900-01-01, endDateが指定されない場合は現在として扱う
// @Summary		指定された期間内に視聴された動画のタグのサマリーを返す
// @Produce		json
// @Param			date		query		string	false	"yyyy-mm-dd"
// @Param			startDate	query		string	false	"yyyy-mm-dd"
// @Param			endDate		query		string	false	"yyyy-mm-dd"
// @Success		200			{array}		data.TagNameCount
// @Failure		400			{object}	ErrorResponse
// @Failure		500			{object}	ErrorResponse
// @Router			/summary [get]
func GetSummary(c *gin.Context) {
	date := c.Query("date")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	var st, et time.Time
	var err error

	if date != "" {
		st, err = time.Parse("2006-01-02", date)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Invalid date format. date format: yyyy-mm-dd",
			})
			return
		}

		st = st.Add(5 * time.Hour)
		et = st.Add(24 * time.Hour)
	} else {
		if startDate == "" {
			st, _ = time.Parse("2006-01-02", "1900-01-01")
		} else {
			st, err = time.Parse("2006-01-02", startDate)
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, ErrorResponse{
					Code:    http.StatusBadRequest,
					Message: "Invalid startDate format. date format: yyyy-mm-dd",
				})
				return
			}
		}

		if endDate == "" {
			et = time.Now()
		} else {
			et, err = time.Parse("2006-01-02", endDate)
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, ErrorResponse{
					Code:    http.StatusBadRequest,
					Message: "Invalid endDate format. date format: yyyy-mm-dd",
				})
				return
			}
		}

		st = st.Add(5 * time.Hour)
		et = et.Add(29 * time.Hour)
	}

	summary, err := logic.CalcSummary(st, et)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, summary)
}
