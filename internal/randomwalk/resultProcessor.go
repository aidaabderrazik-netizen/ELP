package randomwalk

//fonction qui recoit et qui process
func ResultsPro(results <-chan WalkResult, nodeCounts map[int64]int, totalSteps *int64, done chan<- bool) { //done correspond au signal de fin
	for res := range results { //boucle jusqu'à fermeture du channel
		for _, node := range res.Visits { //pour chauqe noeud visité
			nodeCounts[node]++
		}
		*totalSteps += res.Steps
	}
	done <- true
}
