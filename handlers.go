package main

import (
	"net/http"
    "os"
    "time"
    "strconv"
    "log"

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
	
	err = UploadToS3(tf, file.uid)
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