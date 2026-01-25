package randomwalk

import (
	"sort"
)

type NodeProb struct {
	Node int64   `json:"node"`
	Prob float64 `json:"prob"`
}

func TopK(probs map[int64]float64, k int) []NodeProb {
	var list []NodeProb
	for n, p := range probs {
		list = append(list, NodeProb{n, p})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Prob > list[j].Prob
	})

	if len(list) < k {
		return list
	}
	return list[:k]
}
