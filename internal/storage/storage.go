package storage

import (
	"encoding/json"
	"log"
	"sync"
	bolt "go.etcd.io/bbolt"
	"log-ingestor/models"
)

var (
	db       *bolt.DB
	dbOnce   sync.Once
	dbPath   = "logs.db"
	logsBucket = []byte("LogsBucket")
)

func initialize() {
	dbOnce.Do(func() {
		var err error
		db, err = bolt.Open(dbPath, 0600, nil)
		if err != nil {
			log.Fatalf("Failed to open BoltDB: %v", err)
		}

		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists(logsBucket)
			return err
		})

		if err != nil {
			log.Fatalf("Failed to create bucket: %v", err)
		}
	})
}

func AddLog(logEntry models.LogEntry) error {
	initialize()

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(logsBucket)

		// choose a unique key, sm like  auto-incrementing integer sequence by boltdb
		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		key := itob(id)

		value, err := json.Marshal(logEntry)
		if err != nil {
			return err
		}
		return b.Put(key, value)
	})
}

func GetAllLogs()([]models.LogEntry, error) {
	initialize()

	var logs []models.LogEntry
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(logsBucket)
		return b.ForEach(func(k, v []byte) error {
			var logEntry models.LogEntry
			if err := json.Unmarshal(v, &logEntry); err != nil {
				return err
			}
			logs = append(logs, logEntry)
			return nil
		})
	})
	return logs, err
}

// BoltDB stores keys and values as byte slices
func itob(v uint64) []byte {
	b := make([]byte, 8)
	for i := uint(0); i < 8; i++ {
		b[7-i] = byte(v >> (i * 8))
	}
	return b
}