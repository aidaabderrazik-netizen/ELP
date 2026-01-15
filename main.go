package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// var graph = map[int][]int{
// 	0: {1, 3},
// 	1: {0, 2, 4},
// 	2: {1, 5},
// 	3: {0, 4, 6},
// 	4: {1, 3, 5},
// 	5: {2, 4, 8},
// 	6: {3, 7},
// 	7: {6, 8},
// 	8: {5, 7, 9},
// 	9: {8},
// }

var graph map[int64][]int64

func main() {
	rand.Seed(time.Now().UnixNano())

	// Charger le graphe depuis le CSV

	var err error
	graph, err = fonc("lyon_graph.csv")
	if err != nil {
		fmt.Println("Erreur lors du chargement du graphe :", err)
		return
	}

	initNodes()

	// S’assurer que tous les voisins existent comme clé

	for _, voisins := range graph {
		for _, v := range voisins {
			if _, ok := graph[v]; !ok {
				graph[v] = []int64{}
			}
		}
	}

	// Vérification rapide

	// fmt.Println("Noeud 143403 ->", graph[143403])

	//paramètre du test
	numWalks := 100
	duree := 120 * time.Second // 2 minutes (sinon ça prend 1000 ans)
	startNode := int64(143403) // on prend le vrai premier noeud

	fmt.Printf("=== Demo : 1 goroutine qui effectue la marche aléatoire pendant %d secondes ===\n", duree/1000000000)
	//Création du canal pour les results
	results1 := make(chan []int64, numWalks)

	//Création du canal pour savoir quand le process est fini
	done1 := make(chan bool)

	//Compteur global
	nodeCounts1 := make(map[int64]int)

	var wg1 sync.WaitGroup

	go resultsPro(results1, nodeCounts1, done1)
	wg1.Add(1)
	go walkWorker(startNode, duree, results1, &wg1)

	//fermeture du canal
	go func() {
		wg1.Wait()
		close(results1)
	}()
	start := time.Now()
	//Attente du traitement du processeur
	<-done1 //en gros attente de la lecture du channel
	fmt.Println("Temps écoulé :", time.Since(start))

	//Affichage des resultats
	fmt.Println("Nombre de visites par noeud :")
	total1 := 0
	for node1, count1 := range nodeCounts1 {
		fmt.Printf("Noeud %d -> %d visites\n", node1, count1)
		total1 += count1
	}
	fmt.Println("Total visits:", total1)

	//Calculer et afficher les probabilités

	probs1 := ComputeProbabilities(nodeCounts1)

	fmt.Println("\nProbabilités estimées avec 1 goroutine :")
	for node, p := range probs1 {
		fmt.Printf("Noeud %d -> %.10f\n", node, p)
	}

	fmt.Printf("\n=== Demo : %d goroutines qui effectue la marche aléatoire pendant %d secondes ===\n", numWalks, duree/1000000000)
	//Création du canal pour les results
	results := make(chan []int64, numWalks)

	//Création du canal pour savoir quand le process est fini
	done := make(chan bool)

	//Compteur global
	nodeCounts := make(map[int64]int)

	var wg sync.WaitGroup

	go resultsPro(results, nodeCounts, done)
	for i := 0; i < numWalks; i++ {
		wg.Add(1)
		go walkWorker(startNode, duree, results, &wg)

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

	//Affichage des resultats
	fmt.Println("Nombre de visites par noeud :")
	total := 0
	for node, count := range nodeCounts {
		fmt.Printf("Noeud %d -> %d visites\n", node, count)
		total += count
	}
	fmt.Println("Total des visites:", total)

	//Calculer et afficher les probabilités

	probs := ComputeProbabilities(nodeCounts)

	fmt.Printf("\nProbabilités estimées avec %d goroutines: \n", numWalks)
	for node, p := range probs {
		fmt.Printf("Noeud %d -> %.10f\n", node, p)
	}

	// --- Visualisation ---
	layout := forceDirectedLayout(graph, 1200, 1200, 200) // 200 itérations suffisent pour un graphe moyen

	// normaliser les probabilités pour que la couleur soit lisible
	probs1Norm := normalizeProbs(probs1)
	probsNorm := normalizeProbs(probs)

	// dessiner les graphes
	drawGraph("graph_1_goroutine.png", layout, probs1Norm, 1200, 1200)
	drawGraph("graph_multi_goroutines.png", layout, probsNorm, 1200, 1200)

}
