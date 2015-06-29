package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

var region = aws.EUWest

func UploadToS3(file multipart.File, s3Key string) (err error) {

	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal("Incorrect AWS Auth Details ", err)
	}

	connection := s3.New(auth, region)
	bucket := connection.Bucket(os.Getenv("S3_BUCKET"))

	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, file)
	if err != nil {
		log.Fatal("Cannot create file ", err)
	}

	filetype := http.DetectContentType(buffer.Bytes())
	err = bucket.Put(
		s3Key,
		buffer.Bytes(),
		filetype,
		s3.ACL("public-read"),
	)
	return err
}
