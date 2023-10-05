package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var (
	dbConnString = "postgres://user:password@postgres:5432/queue?sslmode=disable"
	queueTable   = "message_queue"
)

func main() {
	http.HandleFunc("/enqueue", enqueueMessageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type RequestBody struct {
	Message string `json:"message"`
}

func enqueueMessageHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var requestBody RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		log.Println("Error parsing request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	message := requestBody.Message

	_, err = db.ExecContext(r.Context(), fmt.Sprintf("INSERT INTO %s (message) VALUES ($1)", queueTable), message)
	if err != nil {
		log.Println("Error enqueueing message:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Println("Message enqueued successfully")
	fmt.Fprintln(w, "Message enqueued successfully")
}
