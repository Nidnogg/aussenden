package main

import (
	//"fmt"
	"log"
	"os"
	"time"
	"path/filepath"
	"html/template"
	"net/http"


)

var port string = "8000"

func loggerHandler(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	t2 := time.Now()
	defer log.Printf("[%s] %q %v", r.Method, r.URL.String(), t2.Sub(t1))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	loggerHandler(w, r);

	// Sets up paths for layout template. filepath.Clean sanitizes User input
	layout_path := filepath.Join("templates", "layout.html")
	file_path := filepath.Join("templates", filepath.Clean(r.URL.Path))
	
	// Returns 404 if template doesnt exist
	info, err := os.Stat(file_path)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Returns 404 if request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	// Serves template
	custom_template, err := template.ParseFiles(layout_path, file_path)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return 
	}


	// Executes template with error clauses
	if err := custom_template.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}	
}

func main() {
	
	// Sets up file root directory
	file_server := http.FileServer(http.Dir("client/public"))
	http.Handle("/client/public/", http.StripPrefix("/client/public/", file_server))
	
	// Index handler servers template
	http.HandleFunc("/", indexHandler) 
	log.Println("Listening on port:" + port)
	log.Fatal(http.ListenAndServe(":" + port, nil)) //returns only if an error occurs
}


