package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(".."))
	mux.Handle(`/YANDEX_PRACTICUM/`, http.StripPrefix(`/YANDEX_PRACTICUM/`, fs))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./main.go")
	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
