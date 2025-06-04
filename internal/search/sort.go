package search

import (
    "log-ingestor/models"
    "sort"
)

func SortLogs(logs []models.LogEntry, field, order string) {
    asc := (order != "desc")
    sort.Slice(logs, func(i, j int) bool {
        switch field {
        case "timestamp":
            return (logs[i].Timestamp < logs[j].Timestamp) == asc
        case "level":
            return (logs[i].Level < logs[j].Level) == asc
        case "resourceId":
            return (logs[i].ResourceID < logs[j].ResourceID) == asc
        case "traceId":
            return (logs[i].TraceID < logs[j].TraceID) == asc
        case "spanId":
            return (logs[i].SpanID < logs[j].SpanID) == asc
        case "parentResourceId":
            return (logs[i].Metadata.ParentResourceID < logs[j].Metadata.ParentResourceID) == asc
        default:
            return true
        }
    })
}
