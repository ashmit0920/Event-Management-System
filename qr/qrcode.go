package qr

import (
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

// generates a QR code for the given data and returns it as a byte slice
func GenerateQRCode(data string) ([]byte, error) {
	qr, err := qrcode.Encode(data, qrcode.Medium, 256)
	if err != nil {
		log.Println("Error generating QR code:", err)
		return nil, err
	}
	return qr, nil
}

// Generates and saves QR to a file
func SaveQRCodeToFile(data string, filename string) error {
	err := qrcode.WriteFile(data, qrcode.Medium, 256, filename)
	if err != nil {
		log.Println("Error saving QR code to file:", err)
		return err
	}
	return nil
}
