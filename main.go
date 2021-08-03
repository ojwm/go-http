package main

import (
	"fmt"
	"log"
	"net/http"
)

func health(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "ok\n")
}

func mp3(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./files/Marsel_Minga_-_01_-_Demonstration.mp3")
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/mp3", mp3)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
