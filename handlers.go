package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const (
	S3URL = "https://s3-eu-west-1.amazonaws.com/"
)

func index(w http.ResponseWriter, r *http.Request) {
	t := GetToken(w, r)
	getView(w, "index", t)
}

func imageUpload(w http.ResponseWriter, r *http.Request) {

	// Get the image from POST data
	f, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()

	token := r.FormValue("token")
	isValid := IsValidToken(token, r)
	if !isValid {
		http.Error(w, "Invalid Token", 500)
		return
	}

	// Calculate a Filename
	// (currently: based on original name and unix time)
	uniqueId := strconv.FormatInt(time.Now().Unix(), 10)
	file := File{
		Filename: uniqueId + "-" + header.Filename,
		uid:      uniqueId,
	}

	err = UploadToS3(f, file.Filename)
	err = nil
	if err != nil {
		log.Fatal("Cannot add to S3 ", err)
	}

	json, err := json.Marshal(file)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", json)
}

func imageView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := File{
		Filename: S3URL + os.Getenv("S3_BUCKET") + "/" + vars["uid"],
	}
	getView(w, "view", f)
}
