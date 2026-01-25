package randomwalk

import (
	"sync"
	"time"
)

type WalkResult struct {
	Visits []int64
	Steps  int64
}

func RunRandomWalks(numWalks int, duration time.Duration) (int64, map[int64]float64) {

	results := make(chan WalkResult, numWalks)
	done := make(chan bool)

	nodeCounts := make(map[int64]int)
	var totalSteps int64 = 0

	var wg sync.WaitGroup

	go ResultsPro(results, nodeCounts, &totalSteps, done)

	startNode := int64(143403) // ou tiré aléatoirement

	for i := 0; i < numWalks; i++ {
		wg.Add(1)
		go WalkWorker(startNode, duration, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	<-done

	probs := ComputeProbabilities(nodeCounts)

	return totalSteps, probs
}
