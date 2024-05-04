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

func SearchVideosByTag(tag string) ([]data.VideoInfoWithoutTags, error) {
	normalizedTag := NormalizeTagName(tag)
	videoInfo, err := data.SearchVideosByTag(normalizedTag)
	if err != nil {
		return nil, err
	}
	return videoInfo, nil
}

func SearchVideosByTitle(title string) ([]data.VideoInfoWithoutTags, error) {
	videoInfo, err := data.SearchVideosByTitle(title)
	if err != nil {
		return nil, err
	}
	return videoInfo, nil
}
