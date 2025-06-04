# Log Ingestor System

A simple web application to **search, view, paginate, sort, and download logs** stored in BoltDB. Built with Go and HTML/CSS.

---

##  Features

- Search logs using filters like:
  - Level
  - Message
  - Resource ID
  - Trace ID
  - Span ID
  - Commit
  - Parent Resource ID
  - Date & Time Range
- Pagination (Next / Previous)
- Sort by any column (ascending/descending)
- Highlight matched message keywords
- CSV download of filtered results
---

## Tech Stack

| Component         | Technology        |
|------------------|-------------------|
| Backend          | Go (Golang)       |
| Database         | BoltDB (embedded key-value store) |
| Templating       | Go's `html/template` |
| Frontend         | HTML5, CSS3       |
| Styling          | Pure CSS (no frameworks) |
| Build Tool       | `go build`        |
| API handling     | Postman           |
---

## Data format
```
{
  "Timestamp": "2025-06-04T10:23:45Z",
  "Level": "INFO",
  "Message": "User login successful",
  "ResourceID": "abc-123",
  "TraceID": "xyz-789",
  "SpanID": "def-456",
  "Commit": "a1b2c3d4",
  "Metadata": {
    "ParentResourceID": "root-999"
  }
}
```

![image](https://github.com/user-attachments/assets/8ad3ca02-98da-42b2-a3d6-b6e20ac1e0c7)

![image](https://github.com/user-attachments/assets/0a34540b-08de-4c44-bce5-8a274ac58f82)

![image](https://github.com/user-attachments/assets/0df95108-b9cb-43a5-81f0-cf564c87f2b9)


