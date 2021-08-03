package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func health(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "ok\n")
}

func mp3(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./files/Marsel_Minga_-_01_-_Demonstration.mp3")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", health)
	router.HandleFunc("/mp3", mp3)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
