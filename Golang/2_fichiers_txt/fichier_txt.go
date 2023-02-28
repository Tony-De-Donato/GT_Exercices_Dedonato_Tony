package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	for {
		fmt.Println("Voici les choix disponibles : \n ")
		fmt.Println("1=Récupérer tout le contenu du fichier.")
		fmt.Println("2=Ajouter du texte dans le fichier.")
		fmt.Println("3=Supprimer tout le contenu du fichier.")
		fmt.Println("4=Remplacer le contenu du fichier.")
		fmt.Println("5=Stop")
		fmt.Println("\nEntrez votre choix :")
		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			lireFichier()
		case 2:
			ajouterAuFichier()
		case 3:
			supprimerContenuFichier()
		case 4:
			supprimerContenuFichier()
			ajouterAuFichier()
		case 5:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func lireFichier() {
	var nomFichier string
	fmt.Println("Entrez le nom du fichier :")
	fmt.Scanln(&nomFichier)                   // Récupération du nom du fichier
	contenu, _ := ioutil.ReadFile(nomFichier) // Lecture du fichier
	fmt.Println(string(contenu))              // Affichage du contenu du fichier
}

func ajouterAuFichier() {
	var nomFichier string
	fmt.Println("Nom du fichier :")
	fmt.Scanln(&nomFichier) // Récupération du nom du fichier
	var texte string
	fmt.Println("Texte à ajouter :")
	fmt.Scanln(&texte)                                                   // Récupération du texte à ajouter
	fichier, _ := os.OpenFile(nomFichier, os.O_APPEND|os.O_WRONLY, 0644) // Ouverture du fichier en mode écriture
	fichier.Close()                                                      // Fermeture du fichier
	fichier.WriteString(texte)                                           // Ajout du texte au fichier
	fmt.Println("Le texte a été ajouté au fichier.")
}

func supprimerContenuFichier() {
	var nomFichier string
	fmt.Println("Entrez le nom du fichier :")
	fmt.Scanln(&nomFichier) // Récupération du nom du fichier
	os.Remove(nomFichier)   // Suppression du contenu du fichier
	fmt.Println("Le fichier a été supprimé.")
}
