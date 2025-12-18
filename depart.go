package main

import (
	"fmt"
	"math/rand"
	"time"
)

var graph = map[int][]int{
	0: {1, 3},
	1: {0, 2, 4},
	2: {1, 5},
	3: {0, 4, 6},
	4: {1, 3, 5},
	5: {2, 4, 8},
	6: {3, 7},
	7: {6, 8},
	8: {5, 7, 9},
	9: {8},
}

// fonction de la marche aléatoire

func randomwalk(start int, pas int) []int {
	var compteur int
	position := start
	hist := []int{position} // trace des positions

	for i := 0; i < pas; i++ {
		voisins := graph[position]
		if len(voisins) == 0 {
			break // plus de voisins, on arrête la marche
		}
		// choisir un voisin aléatoire
		position = voisins[rand.Intn(len(voisins))]
		hist = append(hist, position)
		compteur++
	}
	return hist

}

// fonction main

func main() {
	rand.Seed(time.Now().UnixNano()) // initialisation de l'aléatoire

	start := 0
	pas := 10

	walk := randomwalk(start, pas)
	fmt.Println("Marche aléatoire :", walk)
}
