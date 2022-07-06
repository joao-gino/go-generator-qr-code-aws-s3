package main

import (
	"flag"
	"qrcode/uploadToS3"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

func main() {

	// Create a unique file content
	fileContent := uuid.New().String()

	// Initialize flags to be passed to the script
	flag.StringVar(&fileContent, "fc", fileContent, "File Content")
	flag.Parse()

	var png []byte
	png, error := qrcode.Encode(fileContent, qrcode.Medium, 128)

	if error != nil {
		panic(error)
	}

	uploadToS3.Config("sa-east-1", "gino-qrcode-bucket", png)

}
