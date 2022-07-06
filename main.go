package main

import (
	"qrcode/uploadToS3"

	"github.com/skip2/go-qrcode"
)

func main() {

	var png []byte
	png, error := qrcode.Encode("https://example.org", qrcode.Medium, 256)

	if error != nil {
		panic(error)
	}

	uploadToS3.Config("sa-east-1", "gino-qrcode-bucket", png)

}
