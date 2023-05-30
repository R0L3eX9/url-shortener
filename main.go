package main

import (
	"context"
    "encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
    "strings"

	"github.com/redis/go-redis/v9"
)

type URL struct {
	url         string
}

func generate_url() string {
    // TODO: change rand seed
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


    var url string
    json.NewDecoder(r.Body).Decode(&url)
	shorten_url := generate_url()

    ctx := context.Background()
    client := connect_db()
    err := client.Set(ctx, shorten_url, url , 0).Err()

    if err != nil {
        http.Error(w, "Database is down", http.StatusExpectationFailed)
    }


	fmt.Fprintf(w, "Your url is: %s\n", shorten_url)
	// make sure it has not been used
	// insert into db
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
    client := connect_db()
    id := strings.TrimPrefix(r.URL.Path, "/app/")
    ctx := context.Background()
    shorten_url, err := client.Get(ctx, id).Result()
    if err != nil {
        http.Error(w, "URL not found in the database", http.StatusBadRequest)
    }
    http.Redirect(w, r, shorten_url, http.StatusSeeOther)
	// check if the url is in db
	// redirect if it is
}

func connect_db() *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr: "localhost:8888",
        Password: "",
        DB: 0,
    })
}

func main() {
	// POST - /api/create -> give a url and we generate a random url
	// GET  - /app/<shorten-url> -> sends user to the actual url

	http.HandleFunc("/api/create", createHandler)
    http.HandleFunc("/app/", redirectHandler)

    fmt.Println("Serving on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
