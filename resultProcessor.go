package main

//fonction qui recoit et qui process
func resultsPro(results <-chan []int64, nodeCounts map[int64]int, done chan<- bool) { //done correspond au signal de fin
	for hist := range results { //boucle jusqu'à fermeture du channel
		for _, node := range hist { //pour chauqe noeud visité
			nodeCounts[node]++
		}
	}
	done <- true
}
