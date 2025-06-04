package main

import (
	"log"
	"net/http"
	"log-ingestor/internal/api"
)

func main() {
	http.HandleFunc("/logs", api.IngestHandler)
	http.HandleFunc("/logs/all", api.GetLogsHandler)

	http.HandleFunc("/search", api.SearchFormHandler)
	http.HandleFunc("/results", api.SearchHandler)
	http.HandleFunc("/download", api.DownloadHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":3000", nil))
}