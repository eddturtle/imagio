package main

import (
    "os"
    "log"
    "fmt"
    "time"
    "strconv"
	"net/http"
	"encoding/json"

    "github.com/gorilla/mux"
)

const (
	S3URL = "https://s3-eu-west-1.amazonaws.com/"
)

func Index(w http.ResponseWriter, r *http.Request) {
	getView(w, "index", nil)
}

func ImageUpload(w http.ResponseWriter, r *http.Request) {

	// Get the image from POST data
	f, header, err := r.FormFile("image")
	if err != nil {
		log.Printf("No Image, %s", err)
		return
	}
	defer f.Close()

	// Calculate a Filename 
	// (currently: based on original name and unix time)
	uniqueId := strconv.FormatInt(time.Now().Unix(), 10)
	file := File{
		Filename: uniqueId+"-"+header.Filename, 
		uid: uniqueId,
	}

	err = UploadToS3(f, file.Filename)
	if err != nil {
		log.Fatal("Cannot add to S3 ", err)
	}

	json, err := json.Marshal(file)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Fprintf(w, "%s", json)
}

func ImageView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := File{
		Filename: S3URL + os.Getenv("S3_BUCKET") + "/" + vars["uid"],
	}
	getView(w, "view", f)
}