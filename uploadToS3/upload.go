package uploadToS3

import (
	"bytes"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

func Config(awsRegion string, bucketName string, file []byte) {

	// Initialize flags to be passed to the script
	flag.StringVar(&awsRegion, "r", "sa-east-1", "AWS Region")
	flag.StringVar(&bucketName, "b", "gino-qrcode-bucket", "AWS S3 Bucket to upload to")
	flag.Parse()

	// Basic validation to ensure the flags values are not empty
	if awsRegion == "" || bucketName == "" {
		log.Printf("Required arguments have not been provided.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Create an AWS Session. This will use credentials defined in the environment
	session, err := session.NewSession(&aws.Config{Region: aws.String(awsRegion)})
	if err != nil {
		log.Fatalf("could not initialize new aws session: %v", err)
	}

	// Initialize an s3 client from the session created
	s3Client := s3.New(session)

	// Call the upload to s3 file
	err = uploadFileToS3(s3Client, bucketName, file)
	if err != nil {
		log.Fatalf("could not upload file: %v", err)
	}

}

func uploadFileToS3(
	s3Client *s3.S3,
	bucketName string,
	file []byte,
) error {

	// Create a unique file name
	fileName := uuid.New().String() + ".png"

	// Get the file info
	fileSize := int64(len(file))

	// Put the file object to s3 with the file name
	s3Client.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucketName),
		Key:                  aws.String(fileName),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(file),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(file)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return nil
}
