package main

import (
	"net/http"
    "os"
    "io"
    "time"
    "strconv"
    "log"
    "bufio"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	p := Page{Title:"Home", Body:"Hello"}
	getView(w, "index", p)
}

func ImageUpload(w http.ResponseWriter, r *http.Request) {
	uniqueId := strconv.FormatInt(time.Now().Unix(), 10)
	file := File{Filename:"", uid:uniqueId}

	// Get the image from POST data
	f, _, err := r.FormFile("image")
	if err != nil {
		log.Fatal("Image Missing ", err)
		return
	}
	defer f.Close()

	// Create a temp file
	tf, err := os.Create("uploads/"+file.uid)
	if err != nil {
		log.Fatal("No access", err)
		return
	}
	defer tf.Close()

	// Copy Image -> Temp File
	// _, err = io.Copy(tf, f)
	// if err != nil {
	// 	log.Fatal("No copy", err)
	// 	return
	// }

	// Add More here to Save to DB
	// and have the file on S3.
	AWSAuth := aws.Auth{
		AccessKey: "",
		SecretKey: "",
	}
	region := aws.EUWest

	connection := s3.New(AWSAuth, region)
	bucket := connection.Bucket("imagio")

	fileInfo, _ := tf.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	buffer := bufio.NewReader(tf)
	_, err = buffer.Read(bytes)

	filetype := http.DetectContentType(bytes)
	err = bucket.Put(
		file.uid, 
		bytes, 
		filetype, 
		s3.ACL("public-read"),
	)

	if err != nil {
		log.Fatal("Cannot add to S3 ", err)
	}

	http.Redirect(w, r, "/i/"+file.uid, 302)
}

func ImageView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := "uploads/"+vars["uid"]
	if _, err := os.Stat(path); err == nil {
		f := File{Filename: "/image/"+vars["uid"]}
		getView(w, "view", f)
	}
}

func Image(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, "uploads/"+vars["uid"])
}