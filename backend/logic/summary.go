package logic

import (
	"niconico_recap_backend/data"
	"sort"
	"strings"
	"time"
)

type Summary struct {
	TagCount map[string]int `json:"tag_count"`
}

func NormalizeTagName(tag string) string {
	tag = strings.ToLower(tag)
	return tag
}

func CalcSummary(startDate time.Time, endDate time.Time) ([]data.TagNameCount, error) {
	counter := make(map[string]int)

	tagNameCounts, err := data.GetTagNameCount(startDate, endDate)
	if err != nil {
		return nil, err
	}

	for _, tagNameCount := range tagNameCounts {
		normalizedTagName := NormalizeTagName(tagNameCount.TagName)
		counter[normalizedTagName] += tagNameCount.Count
	}

	summary := make([]data.TagNameCount, 0)
	for tag, count := range counter {
		summary = append(summary, data.TagNameCount{TagName: tag, Count: count})
	}

	// sort by count (desc)
	sort.Slice(summary, func(i, j int) bool {
		return summary[i].Count > summary[j].Count
	})

	return summary, nil
}
