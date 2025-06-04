package models

type Metadata struct {
	ParentResourceID string `json:"parentResourceId"`
}

type LogEntry struct {
	Level      string   `json:"level"`
	Message    string   `json:"message"`
	ResourceID string   `json:"resourceId"`
	Timestamp  string   `json:"timestamp"`
	TraceID    string   `json:"traceId"`
	SpanID     string   `json:"spanId"`
	Commit     string   `json:"commit"`
	Metadata   Metadata `json:"metadata"`
}

type SearchQuery struct {
	Level            string
	Message          string
	ResourceID       string
	StartDate        string
	StartTime        string
	EndDate          string
	EndTime          string
	TraceID          string
	SpanID           string
	Commit           string
	ParentResourceID string
}

// {
//   "level": "error",
//   "message": "Failed to connect to DB",
//   "resourceId": "server-1234",
//   "timestamp": "2023-09-15T08:00:00Z",
//   "traceId": "abc-xyz-123",
//   "spanId": "span-456",
//   "commit": "5e5342f",
//   "metadata": {
//     "parentResourceId": "server-0987"
//   }
// }
