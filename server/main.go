package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Start request")
	defer log.Println("End request")

	select {
	case <-ctx.Done():
		log.Println("Request cancelled")
		http.Error(w, "Request cancelled", http.StatusRequestTimeout)
	case <-time.After(5 * time.Second):
		log.Println("Request completed")
		w.Write([]byte("Request completed"))
	}
}
