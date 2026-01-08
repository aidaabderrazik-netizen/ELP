package main

import (
	"fmt"
	"math/rand"
	"sync"
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

func main() {
	rand.Seed(time.Now().UnixNano())

	//paramètre du test
	numWalks := 5
	steps := 2000
	startNode := 0

	//Création du canal pour les results
	results := make(chan []int, numWalks)

	//Création du canal pour savoir quand le process est fini
	done := make(chan bool)

	//Compteur global
	nodeCounts := make(map[int]int)

	var wg sync.WaitGroup

	go resultsPro(results, nodeCounts, done)
	for i := 0; i < numWalks; i++ {
		wg.Add(1) //Ajoute une goroutines
		go walkWorker(startNode, steps, results, &wg)
	}

	//fermeture du canal
	go func() {
		wg.Wait()
		close(results)
	}()

	//Attente du traitement du processeur
	<-done //en gros attente de la lecture du channel

	//Affichage des resultats
	fmt.Println("Nombre de visites par noeud :")
	for node, count := range nodeCounts {
		fmt.Printf("Noeud %d -> %d visites\n", node, count)
	}

	//Calculer et afficher les probabilités
	probs := ComputeProbabilities(nodeCounts)

	fmt.Println("\nProbabilités estimées :")
	for node, p := range probs {
		fmt.Printf("Noeud %d -> %.4f\n", node, p)
	}

}
