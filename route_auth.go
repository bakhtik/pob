package main

import "net/http"

// GET /login
// Show the login page
func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public.navbar", "login")
}

// GET /signup
// Show the signup page
func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public.navbar", "signup")
}

// POST /singup_account
// Create the user account
func signupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}

}
