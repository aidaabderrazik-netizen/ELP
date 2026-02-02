package main

import (
	"ELP/internal/randomwalk" // Package contenant la logique du random walk
	"fmt"                     // Affichage console
	"math/rand"               // Générateur de nombres aléatoires
	"sync"                    // Synchronisation entre goroutines (WaitGroup)
	"time"                    // Gestion du temps et des durées
)

// global : map[noeud] -> liste des voisins
var graph map[int64][]int64

func main() {
	// Initialisation de la graine aléatoire
	// Permet d'avoir des parcours différents à chaque exécution
	rand.Seed(time.Now().UnixNano())

	// -------- Chargement du graphe depuis le fichier CSV --------
	var err error
	graph, err = randomwalk.ChargerGraphe("graph_small.csv")
	if err != nil {
		fmt.Println("Erreur lors du chargement du graphe :", err)
		return
	}

	// S’assurer que tous les voisins existent comme clé dans la map
	// Évite les accès à des clés inexistantes pendant la marche aléatoire
	for _, voisins := range graph {
		for _, v := range voisins {
			if _, ok := graph[v]; !ok { // si le noeuds existe pas dans le graphe mais qu'il est cité comme voisin
				graph[v] = []int64{} // On lui rajoute une liste vide
			}
		}
	}

	// Injection du graphe dans le package randomwalk
	// Celui-ci sera utilisé par toutes les goroutines
	randomwalk.LoadGraph(graph)

	// -------- Paramètres du test --------

	numWalks := 100            // Nombre de goroutines pour le test parallèle
	duree := 120 * time.Second // Durée fixe du test
	startNode := int64(143403) // Noeud de départ du random walk

	fmt.Printf("Graphe utilisé pour les 2 démos: graph_small.csv \n")

	fmt.Printf(
		"=== Demo : 1 goroutine qui effectue la marche aléatoire pendant %d secondes ===\n",
		duree/1000000000,
	)

	// -------- Test avec 1 seule goroutine --------

	// Canal recevant les résultats de la marche aléatoire
	results1 := make(chan randomwalk.WalkResult, numWalks)

	// Canal utilisé pour signaler la fin du traitement
	done1 := make(chan bool)

	// Map comptant le nombre de visites par noeud
	nodeCounts1 := make(map[int64]int)

	// Compteur du nombre total de pas effectués
	var totalSteps1 int64 = 0

	// WaitGroup pour synchroniser la goroutine de calcul
	var wg1 sync.WaitGroup

	// Goroutine qui agrège les résultats (visites + nombre de pas)
	go randomwalk.ResultsPro(results1, nodeCounts1, &totalSteps1, done1)

	// Lancement d’une seule goroutine de random walk
	wg1.Add(1)
	go randomwalk.WalkWorker(startNode, duree, results1, &wg1)

	// Fermeture du canal une fois la goroutine terminée
	go func() {
		wg1.Wait()
		close(results1)
	}()

	start := time.Now()

	// Attente de la fin de l’agrégation des résultats
	<-done1

	fmt.Println("Temps écoulé :", time.Since(start))
	fmt.Println("Nombre total de pas (1 goroutine) :", totalSteps1)
	fmt.Println("Pas par seconde :", float64(totalSteps1)/duree.Seconds())

	// -------- Test avec plusieurs goroutines --------

	fmt.Printf(
		"\n=== Demo : %d goroutines qui effectue la marche aléatoire pendant %d secondes ===\n",
		numWalks,
		duree/1000000000,
	)

	// Canal recevant les résultats des random walks parallèles
	results := make(chan randomwalk.WalkResult, numWalks)

	// Canal signalant la fin du traitement
	done := make(chan bool)

	// Map des visites par noeud
	nodeCounts := make(map[int64]int)

	// Compteur global du nombre de pas
	var totalSteps int64 = 0

	// WaitGroup pour synchroniser toutes les goroutines
	var wg sync.WaitGroup

	// Goroutine d’agrégation des résultats
	go randomwalk.ResultsPro(results, nodeCounts, &totalSteps, done)

	// Lancement de numWalks goroutines en parallèle
	for i := 0; i < numWalks; i++ {
		wg.Add(1)
		go randomwalk.WalkWorker(startNode, duree, results, &wg)
	}

	// Fermeture du canal quand toutes les goroutines sont terminées
	go func() {
		wg.Wait()
		close(results)
	}()

	staart := time.Now()

	// Attente de la fin du traitement
	<-done

	fmt.Println("Temps écoulé :", time.Since(staart))
	fmt.Println("Nombre total de pas (", numWalks, "goroutines ) :", totalSteps)
	fmt.Println("Pas par seconde :", float64(totalSteps)/duree.Seconds())
}
