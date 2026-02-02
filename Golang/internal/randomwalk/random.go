package randomwalk

import (
	"math/rand"
	"time"
)

// fonction de la marche aléatoire
func Randomwalk(start int64, temps time.Duration) ([]int64, int64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //pour que ce soit aléatoire à chaque tirage

	position := start
	hist := []int64{position} // trace des position
	var steps int64 = 0

	endTime := time.Now().Add(temps) //quand arreter

	for time.Now().Before(endTime) {
		voisins := graph[position]
		time.Sleep(1 * time.Millisecond) // sinon le CPU tourne à fond
		if len(voisins) == 0 {
			position = nodes[r.Intn(len(nodes))] //  on tire un nouveau point de départ
			hist = append(hist, position)
			continue

		}

		position = voisins[r.Intn(len(voisins))]
		steps++
		hist = append(hist, position)
	}

	return hist, steps
}

