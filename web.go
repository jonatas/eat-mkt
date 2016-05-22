package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// to get file's size
type Size interface {
	Size() int64
}

var uploadTemplate = template.Must(template.ParseFiles("upload_form.html"))

func indexHandle(w http.ResponseWriter, r *http.Request) {
	if err := uploadTemplate.Execute(w, nil); err != nil {
		log.Fatal("Execute: ", err.Error())
		return
	}
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(128 << 20) //annotate or not gets the same.
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	if sizeInterface, ok := file.(Size); ok {
		fmt.Fprintf(w, "size is: %d", sizeInterface.Size())
	} else {
		fmt.Printf("can't get %s", sizeInterface)
	}

	t, err := os.Create("./" + "file")
	if err != nil {
		log.Println(err)
		return
	}
	defer t.Close()

	http.Redirect(w, r, "/invite", 303)
	return
}

func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/invite", indexHandle)
	http.HandleFunc("/upload", uploadHandle)
	http.ListenAndServe(":8080", nil)
}
