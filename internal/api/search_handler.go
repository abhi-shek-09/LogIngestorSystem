package api

import (
	"encoding/csv"
	"html/template"
	"log-ingestor/internal/search"
	"log-ingestor/models"
	"math"
	"net/http"
	"strconv"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var query models.SearchQuery
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		query = models.SearchQuery{
			Level:            r.FormValue("level"),
			Message:          r.FormValue("message"),
			ResourceID:       r.FormValue("resourceId"),
			TraceID:          r.FormValue("traceId"),
			SpanID:           r.FormValue("spanId"),
			Commit:           r.FormValue("commit"),
			ParentResourceID: r.FormValue("parentResourceId"),
			StartDate:        r.FormValue("startDate"),
			StartTime:        r.FormValue("startTime"),
			EndDate:          r.FormValue("endDate"),
			EndTime:          r.FormValue("endTime"),
		}
	}

	results, err := search.SearchLogs(query)
	if err != nil {
		http.Error(w, "Search error", http.StatusInternalServerError)
		return
	}
	sortField := r.URL.Query().Get("sort")
	sortOrder := r.URL.Query().Get("order")

	search.SortLogs(results, sortField, sortOrder)

	const pageSize = 10
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if start > len(results) {
		start = len(results)
	}
	if end > len(results) {
		end = len(results)
	}

	pagedResults := results[start:end]

	data := struct {
		Results    []models.LogEntry
		TotalPages int
		Page       int
		HasPrev    bool
		HasNext    bool
		SortField  string
		SortOrder  string
		Query      models.SearchQuery
	}{
		Results:    pagedResults,
		TotalPages: int(math.Ceil(float64(len(results)) / float64(10))),
		Page:       page,
		HasPrev:    page > 1,
		HasNext:    end < len(results),
		SortField:  r.URL.Query().Get("sort"),
		SortOrder:  r.URL.Query().Get("order"),
		Query:      query,
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"toggleSortOrder": func(currentField, currentOrder, thisField string) string {
			if currentField == thisField && currentOrder == "asc" {
				return "desc"
			}
			return "asc"
		},
	}

	tmpl := template.Must(template.New("results.html").Funcs(funcMap).ParseFiles("templates/results.html"))
	tmpl.Execute(w, data)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	query := models.SearchQuery{
		Level:            r.FormValue("level"),
		Message:          r.FormValue("message"),
		ResourceID:       r.FormValue("resourceId"),
		TraceID:          r.FormValue("traceId"),
		SpanID:           r.FormValue("spanId"),
		Commit:           r.FormValue("commit"),
		ParentResourceID: r.FormValue("parentResourceId"),
		StartDate:        r.FormValue("startDate"),
		StartTime:        r.FormValue("startTime"),
		EndDate:          r.FormValue("endDate"),
		EndTime:          r.FormValue("endTime"),
	}

	results, err := search.SearchLogs(query)
	if err != nil {
		http.Error(w, "Search failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment;filename=logs.csv")
	w.Header().Set("Content-Type", "text/csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	writer.Write([]string{
		"Timestamp", "Level", "Message", "ResourceID", "TraceID", "SpanID", "Commit", "ParentResourceID",
	})

	for _, log := range results {
		writer.Write([]string{
			log.Timestamp,
			log.Level,
			log.Message,
			log.ResourceID,
			log.TraceID,
			log.SpanID,
			log.Commit,
			log.Metadata.ParentResourceID,
		})
	}
}
