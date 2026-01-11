package main

import (
	"math/rand"
	"time"
)

// fonction de la marche aléatoire

func randomwalk(start int64, temps time.Duration) []int64 {
	position := start
	hist := []int64{position} // trace des positions

	endTime := time.Now().Add(temps) //quand arreter

	for time.Now().Before(endTime) {
		voisins := graph[position]
		if len(voisins) == 0 {
			break
		}

		position = voisins[rand.Intn(len(voisins))]
		hist = append(hist, position)
	}
	return hist
}
