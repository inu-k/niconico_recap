package logic

import (
	"niconico_recap_backend/data"
	"time"
)

func FetchHistory(startDate time.Time, endDate time.Time) ([]data.DetailHistory, error) {
	history, err := data.GetHistory(startDate, endDate)
	if err != nil {
		return nil, err
	}
	return history, nil
}
