package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/redirect1", redirect1Handler)
	http.HandleFunc("/redirect2", redirect2Handler)
	http.HandleFunc("/redirect3", redirect3Handler)
	http.ListenAndServe(":8080", nil)
}

func redirect1Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8080/redirect1.html", http.StatusSeeOther)
}

func redirect2Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8080/redirect2.html", http.StatusSeeOther)
}

func redirect3Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8080/redirect3.html", http.StatusSeeOther)
}
