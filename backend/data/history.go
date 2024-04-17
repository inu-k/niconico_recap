package data

import (

	// "strings"

	"time"
)

type DetailHistory struct {
	VideoId      string `json:"video_id"`
	WatchDate    string `json:"watch_date"`
	Title        string `json:"title"`
	ThumbnailUrl string `json:"thumbnail_url"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetHistory(startDate time.Time, endDate time.Time) ([]DetailHistory, error) {
	rows, err := Db.Query("select A.video_id, A.watch_date, B.title, B.thumbnail_url from history as A inner join video_basic_info as B on A.video_id=B.video_id where A.watch_date>=$1 and A.watch_date<=$2 order by A.watch_date desc",
		startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05"))

	if err != nil {
		return nil, err
	}

	// make list of DetailHistory
	history := make([]DetailHistory, 0)
	for rows.Next() {
		var d DetailHistory
		err = rows.Scan(&d.VideoId, &d.WatchDate, &d.Title, &d.ThumbnailUrl)
		if err != nil {
			return nil, err
		}
		history = append(history, d)
	}
	rows.Close()
	return history, nil
}
