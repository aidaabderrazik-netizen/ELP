package randomwalk

// ComputeProbabilities calcule la probabilité de visite de chaque noeud
// à partir du nombre total de visites observées pendant les random walks.
func ComputeProbabilities(nodeCounts map[int64]int) map[int64]float64 {
	// Map qui contiendra, pour chaque noeud, sa probabilité estimée
	probs := make(map[int64]float64)

	// Calcul du nombre total de visites sur tous les noeuds
	total := 0
	for _, count := range nodeCounts {
		total += count
	}

	// Pour chaque noeud, on divise son nombre de visites
	// par le nombre total de visites pour obtenir une probabilité
	for node, count := range nodeCounts {
		probs[node] = float64(count) / float64(total)
	}

	// Retourne la distribution de probabilité estimée
	return probs
}
