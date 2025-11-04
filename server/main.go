package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Method: r.Method,
		Path:   r.URL.Path,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(response)
}

func main() {
	port := "8080"
	if _port := os.Getenv("PORT"); _port != "" {
		port = _port
	}
	listenAddr := fmt.Sprintf(":%s", port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
