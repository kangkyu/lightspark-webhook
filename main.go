package main

import (
	"io"
	"log"
	"net/http"
)

const (
	path = "/webhooks"
	port = ":8080"
)

func main() {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		_, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal("unable to read body")
		}
	})
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
