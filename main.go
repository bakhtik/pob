package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("POB WebApp", version(), "started at", config.Address)

	// handle static assets
	files := http.FileServer(http.Dir(config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(config.Address, nil))
}
