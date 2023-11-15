package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"


	_ "github.com/lib/pq"
)

// "os/exec"
// "bytes"
// "regexp"

func main() {
	http.HandleFunc("/enqueue", enqueueMessageHandler)
	http.HandleFunc("/health", func(http.ResponseWriter, *http.Request){})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type RequestBody struct {
	Message string `json:"message"`
}

func enqueueMessageHandler(w http.ResponseWriter, r *http.Request) {
	
	dbConnString := os.Getenv("DB_CONN_STRING")
	queueTable := os.Getenv("QUEUE_TABLE")

	db, err := sql.Open("postgres", dbConnString)
	log.Println("DB Connection string: ", dbConnString)
	log.Println("DB Queue Table: ", queueTable)
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
		_, err = db.ExecContext(r.Context(), fmt.Sprintf("SELECT COUNT(*) FROM %s", queueTable))
		return
	}

	log.Println("Message enqueued successfully")
	fmt.Fprintln(w, "Message enqueued successfully")
}