package data

import (
	"time"
)

type TagNameCount struct {
	TagName string `json:"tag_name"`
	Count   int    `json:"count"`
}

func GetTagNameCount(startDate time.Time, endDate time.Time) ([]TagNameCount, error) {
	rows, err := Db.Query("select B.tag, count(*) from history as A inner join video_tag_info as B on A.video_id=B.video_id where A.watch_date>=$1 and A.watch_date<$2 group by B.tag",
		startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05"))

	if err != nil {
		return nil, err
	}

	ret := make([]TagNameCount, 0)
	for rows.Next() {
		var tmp TagNameCount
		err = rows.Scan(&tmp.TagName, &tmp.Count)
		if err != nil {
			return nil, err
		}
		ret = append(ret, tmp)
	}
	rows.Close()

	return ret, nil
}
