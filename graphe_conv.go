package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nomFichier := "data.txt"

	file, err := os.Open(nomFichier)
	if err != nil {
		fmt.Printf("Erreur : Le fichier %s est introuvable.\n", nomFichier)
		return
	}
	defer file.Close()

	graphe := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ligne := scanner.Text()

		colonnes := strings.Fields(ligne)

		if len(colonnes) >= 2 {
			u, _ := strconv.Atoi(colonnes[0])
			v, _ := strconv.Atoi(colonnes[1])

			graphe[u] = append(graphe[u], v)
		}
	}

	if len(graphe) == 0 {
		fmt.Println("Le graphe est vide. Vérifie le contenu de data.txt")
	} else {
		for sommet, voisins := range graphe {
			fmt.Printf("Sommet %d est lié à : %v\n", sommet, voisins)
		}
	}
}
