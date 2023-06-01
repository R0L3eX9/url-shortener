package api

import (
	"context"
	"encoding/json"
	"fmt"
	math_rand "math/rand"
	"net/http"
	"strings"

	"github.com/redis/go-redis/v9"
)

func generate_url() string {
	const (
		CharSet = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		UrlLen  = 10
	)

	randomURL := make([]byte, UrlLen)
	for i := range randomURL {
		idx := math_rand.Intn(len(CharSet))
		randomURL[i] = CharSet[idx]
	}

	return string(randomURL)
}

type Url struct {
	Url string `json:"url"`
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/create" {
		http.Error(w, "404, not found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	var url Url
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, "Data couldn't be decoded", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client := ConnectDb()
	shortenURL := generate_url()

	_, err = client.Get(ctx, shortenURL).Result()
	for err != redis.Nil {
		shortenURL = generate_url()
		_, err = client.Get(ctx, shortenURL).Result()
	}

	err = client.Set(ctx, shortenURL, url.Url, 0).Err()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Database is down", http.StatusExpectationFailed)
		return
	}

	fmt.Fprintf(w, "Your url is: %s\n", shortenURL)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	client := ConnectDb()
	id := strings.TrimPrefix(r.URL.Path, "/app/")
	ctx := context.Background()

	shortenURL, err := client.Get(ctx, id).Result()
	if err != nil {
		http.Error(w, "URL not found in the database", http.StatusBadRequest)
	}
	http.Redirect(w, r, shortenURL, http.StatusSeeOther)
}

func ConnectDb() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:8888",
		Password: "",
		DB:       0,
	})
}
