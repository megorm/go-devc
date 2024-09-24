package main

import (
	"net/http"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World."))
	})

	http.ListenAndServe(":8080", r)
}
