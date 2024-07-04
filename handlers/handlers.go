package handlers

import (
	"fmt"
	"net/http"

	"event_management/db" // Import the db package
	"event_management/qr" // Import the qr package
)

func GenerateQRHandler(w http.ResponseWriter, r *http.Request) {
	// Handle HTTP request to generate QR code for attendee
	// Example usage of db and qr packages
	qrData := "attendee_id_or_unique_code_here"
	qrBytes, err := qr.GenerateQRCode(qrData)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	// Optionally, save QR code to file
	filename := fmt.Sprintf("%s.png", qrData)
	err = qr.SaveQRCodeToFile(qrData, filename)
	if err != nil {
		http.Error(w, "Failed to save QR code to file", http.StatusInternalServerError)
		return
	}

	// Return success response or render HTML with QR code
	fmt.Fprintf(w, "QR code generated successfully and saved as %s", filename)
}

func ScanQRHandler(w http.ResponseWriter, r *http.Request) {
	// Handle HTTP request to scan QR code and verify attendee
	// Example: Use db package to verify attendee against database
	qrData := r.FormValue("qr_code")

	isValidAttendee := db.VerifyAttendee(qrData)
	if isValidAttendee {
		fmt.Fprintf(w, "Valid attendee")
	} else {
		fmt.Fprintf(w, "Invalid attendee")
	}
}
