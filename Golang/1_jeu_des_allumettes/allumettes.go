package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var n int
	fmt.Println("Combien d'allumettes voulez-vous utiliser ? (4 ou plus)")
	fmt.Scanln(&n)

	if n < 4 {
		fmt.Println("Le nombre d'allumettes doit être au moins 4.")
		return
	}
	var tourJoueur1 bool   // true si c'est au joueur 1 de jouer, false sinon
	if rand.Intn(2) == 0 { // Choisir au hasard qui commence grace a la fonction rand.Intn(2) qui renvoie 0 ou 1
		fmt.Println("Le joueur 1 commence.")
		tourJoueur1 = true
	} else {
		fmt.Println("Le joueur 2 commence.")
		tourJoueur1 = false
	}

	for n > 0 { // Tant qu'il reste des allumettes
		fmt.Println("----------------------------------------------------\n")
		fmt.Printf("Il reste %d allumettes.\n", n)
		var choix int
		if tourJoueur1 { // Demander au joueur actuel de prendre des allumettes
			fmt.Println("Joueur 1, combien d'allumettes voulez-vous prendre ? (1-3)")
		} else {
			fmt.Println("Joueur 2, combien d'allumettes voulez-vous prendre ? (1-3)")
		}
		fmt.Scanln(&choix) // Lire le choix du joueur
		if choix < 1 || choix > 3 || choix > n {
			fmt.Println("La valeur saisie n'est pas valide.")
			continue
		}
		n -= choix                 //on décrémente le nombre d'allumettes choisies de n
		tourJoueur1 = !tourJoueur1 // Changer de joueur
	}
	if tourJoueur1 { // le joueur dont c'est le tour a perdu
		fmt.Println("Le joueur 1 a perdu.")
	} else {
		fmt.Println("Le joueur 2 a perdu.")
	}
}
