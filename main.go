package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const (
	path = "/webhooks"
	port = ":8080"
)

func main() {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal("unable to read body")
		}
		println(string(data))
	})
	err := http.ListenAndServe(":"+getEnv("PORT", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
