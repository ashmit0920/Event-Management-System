package main

import (
	"html/template"
	"log"
	"net/http"

	"event_management/db"       // Import the db package
	"event_management/handlers" // Import the handlers package

	"github.com/joho/godotenv"
)

var (
	templates = template.Must(template.ParseFiles(
		"templates/index.html",
	))
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()
	defer db.CloseDB()

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Serve index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", nil)
	})

	// Set up HTTP server and route handlers
	http.HandleFunc("/scan_qr", handlers.ScanQRHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)

	// Start HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
