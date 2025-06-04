package search

import (
	"log-ingestor/models"
	"log-ingestor/internal/storage"
	"strings"
	"time"
)

func match(value, target string) bool {
	return value == "" || strings.Contains(strings.ToLower(target), strings.ToLower(value))
}

func SearchLogs(query models.SearchQuery) ([]models.LogEntry, error) {
	allLogs, _ := storage.GetAllLogs()

	var results []models.LogEntry

	for _, log := range allLogs {
		if !match(query.Level, log.Level) ||
			!match(query.Message, log.Message) ||
			!match(query.ResourceID, log.ResourceID) ||
			!match(query.TraceID, log.TraceID) ||
			!match(query.SpanID, log.SpanID) ||
			!match(query.Commit, log.Commit) ||
			!match(query.ParentResourceID, log.Metadata.ParentResourceID) {
			continue
		}

		startDate := query.StartDate
		startTime := query.StartTime
		endDate := query.EndDate
		endTime := query.EndTime

		var startTimestamp, endTimestamp string
		if startDate != "" && startTime != "" {
			startTimestamp = startDate + "T" + startTime + "Z"
		}
		if endDate != "" && endTime != "" {
			endTimestamp = endDate + "T" + endTime + "Z"
		}

		if startTimestamp != "" || endTimestamp != "" {
			logTime, err := time.Parse(time.RFC3339, log.Timestamp)
			if err != nil {
				continue
			}

			if startTimestamp != "" {
				startTime, _ := time.Parse(time.RFC3339, startTimestamp)
				if logTime.Before(startTime) {
					continue
				}
			}

			if endTimestamp != "" {
				endTime, _ := time.Parse(time.RFC3339, endTimestamp)
				if logTime.After(endTime) {
					continue
				}
			}
		}
		results = append(results, log)
	}
	return results, nil
}
