package main

import (
	"net/http"
	"time"
)

func initRoutes() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/add-member", AddMemberHandler)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "homepage.html", club)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddMemberHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	name := r.Form.Get("name")
	email := r.Form.Get("email")

	err = club.addMember(&Member{
		Name:             name,
		Email:            email,
		RegistrationDate: time.Now(),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}