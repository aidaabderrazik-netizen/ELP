package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"math/rand"
	"os"
	"sort"
)

type NodeVis struct {
	ID int64
	X  float64
	Y  float64
}

func fixedCircularLayout(nodesIDs []int64) map[int64]NodeVis {
	n := len(nodesIDs)
	layout := make(map[int64]NodeVis, n)

	for i, id := range nodesIDs {
		angle := 2 * math.Pi * float64(i) / float64(n)

		layout[id] = NodeVis{
			ID: id,
			X:  0.5 + 0.45*math.Cos(angle),
			Y:  0.5 + 0.45*math.Sin(angle),
		}
	}
	return layout
}

func nodeIDsFromGraph(graph map[int64][]int64) []int64 {
	ids := make([]int64, 0, len(graph))
	for id := range graph {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	return ids
}

func heatColor(p float64) color.RGBA {
	if p < 0 {
		p = 0
	}
	if p > 1 {
		p = 1
	}

	r := uint8(255 * p)
	b := uint8(255 * (1 - p))
	return color.RGBA{R: r, G: 0, B: b, A: 255}
}

func normalizeProbs(probs map[int64]float64) map[int64]float64 {
	minP := math.MaxFloat64
	maxP := 0.0
	for _, p := range probs {
		if p < minP {
			minP = p
		}
		if p > maxP {
			maxP = p
		}
	}

	norm := make(map[int64]float64, len(probs))
	for id, p := range probs {
		if maxP > minP {
			norm[id] = (p - minP) / (maxP - minP)
		} else {
			norm[id] = 0.5
		}
	}
	return norm
}

func drawGraph(
	filename string,
	layout map[int64]NodeVis,
	probs map[int64]float64,
	width, height int,
) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	for id, node := range layout {
		p := probs[id]
		c := heatColor(p)

		x := int(node.X * float64(width))
		y := int(node.Y * float64(height))

		radius := int(3 + 15*p) // taille ∝ proba

		for dx := -radius; dx <= radius; dx++ {
			for dy := -radius; dy <= radius; dy++ {
				if dx*dx+dy*dy <= radius*radius {
					px := x + dx
					py := y + dy
					if px >= 0 && py >= 0 && px < width && py < height {
						img.Set(px, py, c)
					}
				}
			}
		}
	}

	f, _ := os.Create(filename)
	defer f.Close()
	png.Encode(f, img)
}

// Layout force-directed (Fruchterman-Reingold simplifié)
func forceDirectedLayout(graph map[int64][]int64, width, height int, iterations int) map[int64]NodeVis {
	nodes := make(map[int64]NodeVis)
	n := len(graph)

	// 1️⃣ Initialisation aléatoire
	for id := range graph {
		nodes[id] = NodeVis{
			ID: id,
			X:  rand.Float64() * float64(width),
			Y:  rand.Float64() * float64(height),
		}
	}

	k := math.Sqrt(float64(width*height) / float64(n)) // distance idéale

	for iter := 0; iter < iterations; iter++ {
		// forces de répulsion
		for id1, n1 := range nodes {
			fx, fy := 0.0, 0.0
			for id2, n2 := range nodes {
				if id1 == id2 {
					continue
				}
				dx := n1.X - n2.X
				dy := n1.Y - n2.Y
				dist := math.Sqrt(dx*dx + dy*dy + 0.01)
				rep := k * k / dist
				fx += dx / dist * rep
				fy += dy / dist * rep
			}
			n1.X += 0.1 * fx
			n1.Y += 0.1 * fy
			nodes[id1] = n1
		}

		// forces d’attraction (arêtes)
		for id1, voisins := range graph {
			n1 := nodes[id1]
			for _, id2 := range voisins {
				n2 := nodes[id2]
				dx := n2.X - n1.X
				dy := n2.Y - n1.Y
				dist := math.Sqrt(dx*dx + dy*dy + 0.01)
				attr := dist * dist / k
				n1.X += 0.1 * dx / dist * attr
				n1.Y += 0.1 * dy / dist * attr
			}
			nodes[id1] = n1
		}

		// limiter les positions à l’écran
		for id, n := range nodes {
			if n.X < 0 {
				n.X = 0
			}
			if n.Y < 0 {
				n.Y = 0
			}
			if n.X > float64(width) {
				n.X = float64(width)
			}
			if n.Y > float64(height) {
				n.Y = float64(height)
			}
			nodes[id] = n
		}
	}

	return nodes
}
