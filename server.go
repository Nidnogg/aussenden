package main

import (
	//"fmt"
	"log"
	"net/http"
	"time"
)

var port string = "8000"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	t2 := time.Now()
	defer log.Println("[%s] %q %v", r.Method, r.URL.String(), t2.Sub(t1))
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("Listening on port: %s", port)
	log.Fatal(http.ListenAndServe(":" + port, nil)) //returns only if an error occurs
}


