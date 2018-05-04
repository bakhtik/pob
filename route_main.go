package main

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	if _, err := session(r); err != nil {
		generateHTML(w, nil, "layout", "public.navbar", "index")
	} else {
		generateHTML(w, nil, "layout", "private.navbar", "index")
	}
}
