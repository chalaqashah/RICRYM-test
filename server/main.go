package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go backend!")
}

func main() {

	godotenv.Load()
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8081")
	http.ListenAndServe(":8081", nil)
}
