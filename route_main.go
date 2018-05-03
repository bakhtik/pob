package main

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public.navbar", "index")
}
