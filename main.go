package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type ShortenURL struct {
	OriginalUrl string
	Url         string
}

func generate_url() string {
	const CHAR_SET = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const URL_LEN = 10

	random_url := make([]byte, URL_LEN)
	for i := range random_url {
		idx := rand.Intn(len(CHAR_SET))
		random_url[i] = CHAR_SET[idx]
	}

	return string(random_url)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/create" {
		http.Error(w, "404, not found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	url := generate_url()
	fmt.Fprintf(w, "Your url is: %s\n", url)
	// generate a random url
	// make sure it has not been used
	// insert into db
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Println(w, "TODO")
	// check if the url is in db
	// redirect if it is
}

func main() {
	// POST - /api/create -> give a url and we generate a random url
	// GET  - /api/<shorten-url> -> sends user to the actual url

	http.HandleFunc("/api/create", createHandler)
	http.HandleFunc("/app/", redirectHandler)

    fmt.Println("Serving on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
