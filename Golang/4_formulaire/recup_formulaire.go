package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "le_formulaire.html")
	})

	http.HandleFunc("/envoie", func(w http.ResponseWriter, r *http.Request) {
		nom := r.FormValue("name")
		email := r.FormValue("email")
		fmt.Println("Nom : ", nom)
		fmt.Println("Email : ", email)
	})

	http.ListenAndServe(":8080", nil)
}
