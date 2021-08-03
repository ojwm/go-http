package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func health(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(map[string]bool{"ok": true})
}

func mp3(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./files/Marsel_Minga_-_01_-_Demonstration.mp3")
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/api/health", health)
	router.HandleFunc("/mp3", mp3)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9000",
		// Good practice to set timeouts to mitigate Slowloris attacks
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Run server in a goroutine so it does not block
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	// Accept graceful shutdowns when quit via SIGINT (Ctrl+C),
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until signal received
	<-c

	// Deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections but will otherwise wait
	// until the timeout deadline
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
