package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	log.Println("starting server...")
        err := http.ListenAndServe(":3000", nil)
        if err != nil {
            log.Fatal(err)
        }
}
