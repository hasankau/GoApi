package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Define a struct to represent the data you want to send and receive via the API.
type Message struct {
	Text string `json:"text"`
}

// Define an HTTP handler function to handle incoming requests.
func handler(w http.ResponseWriter, r *http.Request) {

	// Set the response header to indicate JSON content.
	w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/viwaha_app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Attempt to connect to the database.
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	//fmt.Println("Connected to the database!")

	rows, err := db.Query("SELECT email FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	emails := make(map[string]string)

	// Loop through the rows and add the names to the emails map.
	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			panic(err.Error())
		}
		emails[name] = "" // Assign an empty string as the default value for each name
	}

	// Encode the Message struct to JSON and write it as the response.
	json.NewEncoder(w).Encode(emails)

}

func main() {
	// Define the route and the handler function for the API endpoint.
	http.HandleFunc("/api", handler)

	// Start the HTTP server on port 8080.
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
