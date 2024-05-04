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

type VideoInfoWithoutTags struct {
	VideoId      string `json:"video_id"`
	Title        string `json:"title"`
	ThumbnailUrl string `json:"thumbnail_url"`
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

// search videos by tag
func SearchVideosByTag(tag string) ([]VideoInfoWithoutTags, error) {
	rows, err := Db.Query("select A.video_id, A.title, A.thumbnail_url from video_basic_info as A inner join video_tag_info as B on A.video_id=B.video_id where lower(B.tag)=$1", tag)
	if err != nil {
		return nil, err
	}

	ret := make([]VideoInfoWithoutTags, 0)
	for rows.Next() {
		var tmp VideoInfoWithoutTags
		err = rows.Scan(&tmp.VideoId, &tmp.Title, &tmp.ThumbnailUrl)
		if err != nil {
			return nil, err
		}
		ret = append(ret, tmp)
	}
	rows.Close()

	return ret, nil
}

// serach videos by title
func SearchVideosByTitle(title string) ([]VideoInfoWithoutTags, error) {
	rows, err := Db.Query("select video_id, title, thumbnail_url from video_basic_info where title like $1", "%"+title+"%")
	if err != nil {
		return nil, err
	}

	ret := make([]VideoInfoWithoutTags, 0)
	for rows.Next() {
		var tmp VideoInfoWithoutTags
		err = rows.Scan(&tmp.VideoId, &tmp.Title, &tmp.ThumbnailUrl)
		if err != nil {
			return nil, err
		}
		ret = append(ret, tmp)
	}
	rows.Close()

	return ret, nil
}
