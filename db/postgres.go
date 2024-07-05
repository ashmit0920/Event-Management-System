package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

// initialize the database connection
func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to PostgreSQL")
}

func CloseDB() {
	db.Close()
}

func InsertAttendee(name, email, roll string) (int, error) {
	var id int
	query := "INSERT INTO attendees (name, email, roll) VALUES ($1, $2, $3) RETURNING id"

	err := db.QueryRow(query, name, email, roll).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func SaveQRData(attendee_id int, qrData string) error {
	query := "UPDATE attendees SET qr_data = $1 WHERE id = $2"
	_, err := db.Exec(query, qrData, attendee_id)

	if err != nil {
		return err
	}
	return nil
}

func VerifyAttendee(qrData string) bool {
	var count int
	query := "SELECT COUNT(*) FROM attendees WHERE qr_data = $1"

	err := db.QueryRow(query, qrData).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}
