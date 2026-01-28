package randomwalk

/// Fonction qui lance la marche en goroutine : appelle randomwalk, envoie le resultat dans un channel, et marque la fin
import (
	"sync"
	"time"
)

func WalkWorker(
	start int64,
	temps time.Duration,
	results chan<- WalkResult, //envoie dans le canal
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	hist, steps := Randomwalk(start, temps)
	results <- WalkResult{
		Visits: hist,
		Steps:  steps,
	} //envoie de l'historique
}
