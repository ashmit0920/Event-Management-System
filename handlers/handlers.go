package handlers

import (
	"fmt"
	"net/http"

	"event_management/db" // Import the db package
	"event_management/qr" // Import the qr package
)

func ScanQRHandler(w http.ResponseWriter, r *http.Request) {
	// Handle HTTP request to scan QR code and verify attendee
	qrData := r.FormValue("qr_code")

	isValidAttendee := db.VerifyAttendee(qrData)
	if isValidAttendee {
		fmt.Fprintf(w, "Valid attendee")
	} else {
		fmt.Fprintf(w, "Invalid attendee")
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	name := r.FormValue("name")
	email := r.FormValue("email")
	roll := r.FormValue("roll")

	// Store attendee data in the database
	attendeeID, err := db.InsertAttendee(name, email, roll)
	if err != nil {
		http.Error(w, "Failed to register attendee", http.StatusInternalServerError)
		return
	}

	// Generate QR code for the attendee
	qrData := fmt.Sprintf("attendee_%d", attendeeID)
	err = qr.SaveQRCodeToFile(qrData, fmt.Sprintf("static/qr_%d.png", attendeeID))
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	// Saving qrData into the database
	err = db.SaveQRData(attendeeID, qrData)
	if err != nil {
		http.Error(w, "Failed to save QR data", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Registration successful! QR code generated.")
}
