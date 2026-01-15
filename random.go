package main

import (
	"math/rand"
	"time"
)

// listes des noeuds du graphe

var nodes []int64

func initNodes() {
	nodes = make([]int64, 0, len(graph))
	for n := range graph {
		nodes = append(nodes, n)
	}
}

// fonction de la marche aléatoire

func randomwalk(start int64, temps time.Duration) []int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	position := start
	hist := []int64{position} // trace des positions

	endTime := time.Now().Add(temps) //quand arreter

	for time.Now().Before(endTime) {
		voisins := graph[position]
		time.Sleep(1 * time.Millisecond) // sinon le CPU tourne à fond → risqué
		if len(voisins) == 0 {
			position = nodes[r.Intn(len(nodes))] // cul de sac → on tire un nouveau point de départ
			hist = append(hist, position)
			continue

		}

		position = voisins[r.Intn(len(voisins))]
		hist = append(hist, position)
	}
	return hist
}
