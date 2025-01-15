package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "wishal2311"
	dbname   = "wira"
)

func main() {
	acc := NewAccount(1, "JohnDoe", "johndoe@example.com")
	char := NewCharacter(101, 1, 5)
	score := NewScores(1001, 101, 88)

	acc.Display()
	char.Display()
	score.Display()

	// Create connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	// Perform a query
	rows, err := db.Query("SELECT id, username FROM users")
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()
	// Process rows...
	// Iterate through the results
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
