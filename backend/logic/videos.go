package logic

import (
	"niconico_recap_backend/data"
)

func FetchVideoInfo(videoId string) (data.VideoInfo, error) {
	videoInfo, err := data.GetVideoInfo(videoId)
	if err != nil {
		return data.VideoInfo{}, err
	}
	return videoInfo, nil
}
