package main

import (
	"fmt"
	"laporan_keuangan/database"
	"log"
)

func main() {
	// Initialize connection
	db := database.GetConnection()
	defer db.Close()

	// Verify connection
	err := db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Successfully connected to database 'laporan_keuangan'!")

	// Check if we can query users (just for testing)
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		fmt.Printf("Note: Could not query 'users' table (it might be empty or missing): %v\n", err)
	} else {
		fmt.Printf("Found %d users in the database.\n", count)
	}
}
