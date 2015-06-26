package main

import (
	"net/http"
    "os"
    "io"
    "time"
    "strconv"
    "log"
)

func Index(w http.ResponseWriter, r *http.Request) {
	p := Page{Title:"Home", Body:"Hello"}
	getView(w, "index", p)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

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
	_, err = io.Copy(tf, f)
	if err != nil {
		log.Fatal("No copy", err)
		return
	}

	// Add More here to Save to DB
	// and have the file on S3.

	http.Redirect(w, r, "/view?id="+file.uid, 302)
}

func View(w http.ResponseWriter, r *http.Request) {
	path := "uploads/"+r.FormValue("id")
	if _, err := os.Stat(path); err == nil {
		// i := Image{Filename: "/i?id="+r.FormValue("id")}
		f := File{Filename: "/i?id="+r.FormValue("id")}
		getView(w, "view", f)
	}
}