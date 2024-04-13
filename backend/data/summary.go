package data

import (
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Summary struct {
	TagCount map[string]int `json:"tag_count"`
}

type TagNameCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func NormalizeTag(tag string) string {
	tag = strings.ToLower(tag)
	return tag
}

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
// @Success		200			{array}		TagNameCount
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
		if startDate == "" && endDate == "" {
			c.IndentedJSON(http.StatusBadRequest, ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "date or startDate or endDate must be specified",
			})
			return
		}

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

	rows, err := Db.Query("select B.tag, count(*) from history as A inner join video_tag_info as B on A.video_id=B.video_id where A.watch_date>=$1 and A.watch_date<$2 group by B.tag",
		st.Format("2006-01-02 15:04:05"), et.Format("2006-01-02 15:04:05"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	counter := make(map[string]int)

	for rows.Next() {
		var tag string
		var count int
		err = rows.Scan(&tag, &count)
		tag = NormalizeTag(tag)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
		counter[tag] += count
	}
	rows.Close()

	summary := make([]TagNameCount, 0)
	for tag, count := range counter {
		summary = append(summary, TagNameCount{tag, count})
	}

	// sort by count (desc)
	sort.Slice(summary, func(i, j int) bool {
		return summary[i].Count > summary[j].Count
	})

	c.IndentedJSON(http.StatusOK, summary)
}
