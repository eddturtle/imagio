package main

import (
    "os"
    "io"
    "log"
    // "fmt"
    "time"
    "strconv"
	"net/http"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	p := Page{Title:"Home", Body:"Hello"}
	getView(w, "index", p)
}

func ImageUpload(w http.ResponseWriter, r *http.Request) {
	uniqueId := strconv.FormatInt(time.Now().Unix(), 10)

	// Get the image from POST data
	f, header, err := r.FormFile("image")

	file := File{Filename:"", uid:uniqueId+"-"+header.Filename}

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
 	_, err = io.Copy(tf, f)
 	if err != nil {
 		log.Fatal("No copy", err)
 		return
 	}

 	real, err := os.Open("uploads/"+file.uid)

	err = UploadToS3(real, file.uid)
	if err != nil {
		log.Fatal("Cannot add to S3 ", err)
	}

	http.Redirect(w, r, "/i/"+file.uid, 302)
}

func ImageView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := File{Filename: "https://s3-eu-west-1.amazonaws.com/imagio/"+vars["uid"]}
	getView(w, "view", f)
}

// func Image(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	w.Header().Set("Content-Type", "image")

// 	data, err := GetFromS3(vars["uid"])
// 	if err != nil {
// 		log.Fatal("Cannot get from S3 ", err)
// 	}

// 	fmt.Fprintf(w, string(data))
	// http.ServeFile(w, r, "uploads/"+vars["uid"])
// }	