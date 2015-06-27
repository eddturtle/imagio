package main

import (
	"os"
	"bufio"
	"net/http"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

var auth = aws.EnvAuth()
var region = aws.EUWest

func UploadToS3(file *os.File, s3Key string) (err error) {

	connection := s3.New(auth, region)
	bucket := connection.Bucket(os.Getenv("S3_BUCKET"))

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)
	if err != nil {
		log.Fatal("Cannot read from Bytes ", err)
		return
	}

	filetype := http.DetectContentType(bytes)
	err = bucket.Put(
		s3Key, 
		bytes, 
		filetype, 
		s3.ACL("private"),
	)
	return err
}