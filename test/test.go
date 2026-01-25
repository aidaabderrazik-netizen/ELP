package main

import (
	"ELP/internal/randomwalk"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var graph map[int64][]int64

func main() {
	rand.Seed(time.Now().UnixNano()) // permet le tirage aléatoire

	// Charger le graphe depuis le CSV

	var err error
	graph, err = randomwalk.ChargerGraphe("lyon_graph.csv")
	if err != nil {
		fmt.Println("Erreur lors du chargement du graphe :", err)
		return
	}

	// S’assurer que tous les voisins existent comme clé

	for _, voisins := range graph {
		for _, v := range voisins {
			if _, ok := graph[v]; !ok {
				graph[v] = []int64{}
			}
		}
	}
	randomwalk.LoadGraph(graph)

	//paramètre du test
	numWalks := 100
	duree := 120 * time.Second // 2 minutes (sinon ça prend 1000 ans)
	startNode := int64(143403) // on prend le vrai premier noeud

	fmt.Printf("=== Demo : 1 goroutine qui effectue la marche aléatoire pendant %d secondes ===\n", duree/1000000000)
	//Création du canal pour les results
	results1 := make(chan randomwalk.WalkResult, numWalks)

	//Création du canal pour savoir quand le process est fini
	done1 := make(chan bool)

	//Compteur global
	nodeCounts1 := make(map[int64]int)
	var totalSteps1 int64 = 0

	var wg1 sync.WaitGroup

	go randomwalk.ResultsPro(results1, nodeCounts1, &totalSteps1, done1)
	wg1.Add(1) //Ajout d'une goroutine
	go randomwalk.WalkWorker(startNode, duree, results1, &wg1)

	//fermeture du canal
	go func() {
		wg1.Wait()
		close(results1)
	}()
	start := time.Now()
	//Attente du traitement du processeur
	<-done1 //en gros attente de la lecture du channel
	fmt.Println("Temps écoulé :", time.Since(start))
	fmt.Println("Nombre total de pas (1 goroutine) :", totalSteps1)
	fmt.Println("Pas par seconde :", float64(totalSteps1)/duree.Seconds())

	fmt.Printf("\n=== Demo : %d goroutines qui effectue la marche aléatoire pendant %d secondes ===\n", numWalks, duree/1000000000)
	//Création du canal pour les results
	results := make(chan randomwalk.WalkResult, numWalks)

	//Création du canal pour savoir quand le process est fini
	done := make(chan bool)

	//Compteur global
	nodeCounts := make(map[int64]int)
	var totalSteps int64 = 0
	var wg sync.WaitGroup

	go randomwalk.ResultsPro(results, nodeCounts, &totalSteps, done)
	for i := 0; i < numWalks; i++ {
		wg.Add(1)
		go randomwalk.WalkWorker(startNode, duree, results, &wg)

	}
	//fermeture du canal
	go func() {
		wg.Wait()
		close(results)
	}()
	staart := time.Now()
	//Attente du traitement du processeur
	<-done //en gros attente de la lecture du channel
	fmt.Println("Temps écoulé :", time.Since(staart))
	fmt.Println("Nombre total de pas (", numWalks, "goroutines ) :", totalSteps)
	fmt.Println("Pas par seconde :", float64(totalSteps)/duree.Seconds())

}
