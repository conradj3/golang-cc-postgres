package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	dbConnString        = "postgres://user:password@postgres:5432/queue?sslmode=disable"
	queueTable          = "message_queue"
	desiredMessageCount = 5
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	log.Println("Consumer started")
	log.Println("Listening for messages in the queue...")

	var processedMessages int

	for processedMessages < desiredMessageCount {
		message, err := dequeueMessage(ctx, db)
		if err != nil {
			if err != sql.ErrNoRows {
				log.Println("Error dequeuing message:", err)
			}
		} else {
			log.Println("Processing message:", message)
			processedMessages++
			// Simulate message processing
			if processMessageErr := processMessage(message); processMessageErr != nil {
				log.Println("Error processing message:", processMessageErr)
			} else {
				log.Println("Message processed successfully")
			}
		}
	}

	log.Println("Consumer finished")
}

func dequeueMessage(ctx context.Context, db *sql.DB) (string, error) {
	var message string

	err := db.QueryRowContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id = (SELECT id FROM %s ORDER BY timestamp LIMIT 1) RETURNING message", queueTable, queueTable)).Scan(&message)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Error dequeuing message:", err)
		}
		return "", err
	}

	return message, nil
}

func processMessage(message string) error {
	// Simulated message processing
	time.Sleep(2 * time.Second)
	// Simulated processing error
	if message == "error" {
		return fmt.Errorf("simulated processing error")
	}
	return nil
}
