package main

func ComputeProbabilities(
	nodeCounts map[int]int,
) map[int]float64 {
	probs := make(map[int]float64)

	total := 0
	for _, count := range nodeCounts {
		total += count
	}
	for node, count := range nodeCounts {
		probs[node] = float64(count) / float64(total)
	}
	return probs
}
