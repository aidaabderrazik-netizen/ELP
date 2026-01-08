package main

/// Fonction qui lance la marche en goroutine : appelle randomwalk, envoie le resultat dans un channel, et marque la fin
import (
	"sync"
)

func walkWorker(
	start int,
	pas int,
	result chan<- []int, //envoie dans le canal
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	hist := randomwalk(start, pas)
	result <- hist //envoie de l'historique
}
