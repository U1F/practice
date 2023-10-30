// Assuming this file is in a directory named 'datab', making it part of the 'datab' package.
package datab

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// so much complaint about a library that golang claims exists and it can use..... Ok
	_ "github.com/lib/pq"
)

// DB is a global database connection pool.
var DB *sql.DB

// InitializeDB sets up the database connection pool.
func InitializeDB(host, port, user, password, dbname string) {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection is successful
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database.")
}

// Handler function for a hypothetical endpoint that the frontend might interact with.
func MyEndpointHandler(w http.ResponseWriter, r *http.Request) {
	// Hypothetical query: SELECT * FROM my_table;
	rows, err := DB.Query("SELECT * FROM my_table")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Here you would process the rows and return the appropriate response.
	// This is just a simplified placeholder.
	for rows.Next() {
		// Scan rows and do something with the data.
		// Typically you would send back a JSON to your frontend.
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write a success response back to the requester (e.g., frontend).
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Query executed successfully!")
}
