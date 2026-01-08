package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// graph test

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

func worker(id int, wg *sync.WaitGroup) { // définition de la fonction worker, id : identifiant de la goroutine
	defer wg.Done()                // quand le worker est terminé, on signale au WaitGroup
	start := rand.Intn(len(graph)) // sommet de départ random : chaque goroutine commence à un sommet différent
	pas := 20                      // nombre de pas

	hist := randomwalk(start, pas) // hist : historique des positions

	fmt.Println("Goroutine", id, ":", hist) // affichage de l'historique des positions pour chaque goroutine
}

func main() {
	var wg sync.WaitGroup // création d'un WaitGroup pour synchroniser les goroutines

	nbGoroutines := 10 // nombre de goroutines à lancer

	for i := 0; i < nbGoroutines; i++ {
		wg.Add(1)         // on ajoute une goroutine au WaitGroup
		go worker(i, &wg) // lancement de la goroutine
	}

	wg.Wait() // on attend que toutes les goroutines soient terminées
}

// // fonction main

// func main() {
// 	rand.Seed(time.Now().UnixNano()) // initialisation de l'aléatoire

// 	start := 0
// 	pas := 10

// 	walk := randomwalk(start, pas)
// 	fmt.Println("Marche aléatoire :", walk)
// }
