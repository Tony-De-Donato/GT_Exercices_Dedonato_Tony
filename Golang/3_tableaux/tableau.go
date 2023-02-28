package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func genererTableau(n int) [][]int {
	// j'ai considéré que le tableau est un tableau à deux dimensions de taille n x n, car il est trop complexe d'afficher un vrai tableau à plus de deux dimensions dans un navigateur web

	tableau := make([][]int, n) // Créer un tableau à deux dimensions de taille n x n
	for i := range tableau {
		tableau[i] = make([]int, n) // Créer un tableau à une dimension de taille n
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tableau[i][j] = rand.Intn(10) // Remplir le tableau de valeurs aléatoires inférieures à 10 pour que l'affichage soit plus lisible
		}
	}
	return tableau
}

func afficherTableauHTML(w http.ResponseWriter, tableau [][]int) {

	fmt.Fprint(w, "<html><body><table>") // Début du code HTML

	for _, ligne := range tableau { // Parcourir le tableau
		fmt.Fprint(w, "<tr>")           // Créer une nouvelle ligne
		for _, element := range ligne { // Parcourir la ligne
			fmt.Fprint(w, "<td>", element, "</td>") // Créer une nouvelle cellule
		}
		fmt.Fprint(w, "</tr>") // Fermer la ligne
	}

	fmt.Fprint(w, "</table></body></html>") // Fin du code HTML
}

func main() {
	n := rand.Intn(97) + 3 // Générer un nombre aléatoire entre 3 et 100
	tableau := genererTableau(n)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // Afficher le tableau dans le navigateur web
		afficherTableauHTML(w, tableau)
	})
	http.ListenAndServe(":"+strconv.Itoa(8080), nil) // Lancer le serveur web
}
