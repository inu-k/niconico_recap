package data

import (
	"errors"
)

type VideoInfo struct {
	VideoId      string   `json:"video_id"`
	Title        string   `json:"title"`
	Tags         []string `json:"tags"`
	ThumbnailUrl string   `json:"thumbnail_url"`
}

func GetVideoInfo(videoId string) (VideoInfo, error) {
	rows, err := Db.Query("select video_id, title, thumbnail_url from video_basic_info where video_id=$1", videoId)

	if err != nil {
		return VideoInfo{}, err
	}

	var v VideoInfo
	if rows.Next() {
		err = rows.Scan(&v.VideoId, &v.Title, &v.ThumbnailUrl)
		if err != nil {
			return VideoInfo{}, err
		}
	} else {
		return VideoInfo{}, errors.New("video not found")
	}
	rows.Close()

	tags := make([]string, 0)

	rows, err = Db.Query("select tag from video_tag_info where video_id=$1", videoId)
	if err != nil {
		return VideoInfo{}, err
	}

	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			return VideoInfo{}, err
		}
		tags = append(tags, tag)
	}

	v.Tags = tags
	return v, nil
}
