package datab

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/a-h/templ"
	_ "github.com/lib/pq"
)

// DBParams contains parameters for database connection.
type DBParams struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func main() {
	if len(os.Args) < 6 {
		log.Fatalln("Usage: go run createDBWithTempl.go <host> <port> <user> <password> <dbname>")
	}

	// Load the configuration template file.
	templateData, err := os.ReadFile("dbconfig.templ")
	if err != nil {
		log.Fatalf("Unable to read dbconfig.templ: %v", err)
	}

	// Create an instance of DBParams with data from the command line arguments.
	params := DBParams{
		Host:     os.Args[1],
		Port:     os.Args[2],
		User:     os.Args[3],
		Password: os.Args[4],
		DBName:   os.Args[5], // This will not be used in the connection string but for the creation command
	}

	// Parse the template and replace placeholders with actual values.
	templateString := string(templateData)
	connectionString, err := templ.New().Execute(templateString, params)
	if err != nil {
		log.Fatalf("Error processing template: %v", err)
	}

	// Connect to the PostgreSQL server.
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to the Postgres server: %v", err)
	}
	defer db.Close()

	// Verify connection.
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to establish a connection: %v", err)
	}
	fmt.Println("Successfully connected!")

	// Create the database. We're connecting to the default database, then issuing the create command.
	_, err = db.Exec("CREATE DATABASE " + params.DBName)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	fmt.Printf("Database %s successfully created.\n", params.DBName)
}
