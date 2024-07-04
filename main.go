package main

import (
	"net/http"

	"event_management/db"       // Import the db package
	"event_management/handlers" // Import the handlers package
)

func main() {
	// Connect to PostgreSQL database
	db.ConnectDB()
	defer db.CloseDB()

	// Set up HTTP server and route handlers
	http.HandleFunc("/generate_qr", handlers.GenerateQRHandler)
	http.HandleFunc("/scan_qr", handlers.ScanQRHandler)

	// Start HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
