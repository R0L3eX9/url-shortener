package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	math_rand "math/rand"
	"net/http"

	"github.com/R0L3eX9/url-shortener/api"
)

func init() {
	var sd [8]byte
	_, err := crypto_rand.Read(sd[:])
	if err != nil {
		panic("Cannot seed math/rand with crypto/rand")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(sd[:])))
}

func main() {
	http.HandleFunc("/api/create", api.CreateHandler)
	http.HandleFunc("/app/", api.RedirectHandler)

	fmt.Println("Serving on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
