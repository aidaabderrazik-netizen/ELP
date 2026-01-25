package randomwalk

func ComputeProbabilities(nodeCounts map[int64]int) map[int64]float64 {
	probs := make(map[int64]float64)

	total := 0
	for _, count := range nodeCounts {
		total += count
	}
	for node, count := range nodeCounts {
		probs[node] = float64(count) / float64(total)
	}
	return probs
}
