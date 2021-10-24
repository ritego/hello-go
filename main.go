package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Say Hello World in console
	fmt.Println("Hello World")

	// make the port number a reusable variable
	addr := ":7777"

	// define routes
	router := http.DefaultServeMux
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World"))
	})

	// configure server
	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// start accepting requests
	log.Printf("Server running on: %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
