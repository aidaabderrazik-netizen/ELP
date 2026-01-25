package randomwalk

import (
	"sync"
	"time"
)

// WalkResult représente le résultat produit par une goroutine de random walk
// - Visits : liste des nœuds visités pendant la marche
// - Steps  : nombre total de pas effectués pendant la durée imposée
type WalkResult struct {
	Visits []int64
	Steps  int64
}

// RunRandomWalks lance plusieurs random walks en parallèle et agrège les résultats
// - numWalks : nombre de goroutines (marches aléatoires) à lancer
// - duration : durée pendant laquelle chaque marche s’exécute
// Retourne :
// - le nombre total de pas effectués
// - la distribution de probabilité des visites par nœud
func RunRandomWalks(numWalks int, duration time.Duration) (int64, map[int64]float64) {

	// Canal qui reçoit les résultats produits par chaque WalkWorker
	results := make(chan WalkResult, numWalks)

	// Canal utilisé pour signaler que l’agrégation des résultats est terminée
	done := make(chan bool)

	// Map globale comptant le nombre de visites par nœud
	nodeCounts := make(map[int64]int)

	// Compteur global du nombre total de pas effectués
	var totalSteps int64 = 0

	// WaitGroup pour attendre la fin de toutes les goroutines WalkWorker
	var wg sync.WaitGroup

	// Goroutine responsable de :
	// - lire les résultats envoyés dans le channel results
	// - mettre à jour nodeCounts
	// - incrémenter totalSteps
	// - notifier la fin via le channel done
	go ResultsPro(results, nodeCounts, &totalSteps, done)

	// Nœud de départ des random walks
	// Ici il est fixé, mais pourrait être choisi aléatoirement
	startNode := int64(143403)

	// Lancement des goroutines de random walk
	for i := 0; i < numWalks; i++ {
		wg.Add(1)
		go WalkWorker(startNode, duration, results, &wg)
	}

	// Goroutine chargée de fermer le channel results
	// une fois que toutes les goroutines WalkWorker sont terminées
	go func() {
		wg.Wait()
		close(results)
	}()

	// Attente que ResultsPro ait terminé de traiter tous les résultats
	<-done

	probs := ComputeProbabilities(nodeCounts)

	return totalSteps, probs
}
