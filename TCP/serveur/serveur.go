package main

import (
	"ELP/internal/protocol"
	"ELP/internal/randomwalk"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

var maxClients = make(chan struct{}, 4) // limite à 4 clients
var graph map[int64][]int64

func main() {

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Serveur TCP en écoute sur le port 9000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	maxClients <- struct{}{} // semaphore
	defer func() { <-maxClients }()
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	var req protocol.Request
	if err := decoder.Decode(&req); err != nil {
		return
	}
	fmt.Println("Requête reçue :", req)

	// Charger le graphe UNE FOIS
	var err error

	filename := "lyon_graph.csv"

	if req.Graph == "small" {
		filename = "graph_small.csv"
	}

	graph, err := randomwalk.ChargerGraphe(filename)
	if err != nil {
		panic(err)
	}
	randomwalk.LoadGraph(graph)
	// // S’assurer que tous les voisins existent comme clé

	for _, voisins := range graph {
		for _, v := range voisins {
			if _, ok := graph[v]; !ok {
				graph[v] = []int64{}
			}
		}
	}

	// ---- Calcul random walks ----
	duration := time.Duration(req.DurationSec) * time.Second

	steps1, _ := randomwalk.RunRandomWalks(1, duration)
	stepsN, probsN := randomwalk.RunRandomWalks(req.NumWalks, duration)

	topN := randomwalk.TopK(probsN, 5)

	speedup := float64(stepsN) / float64(steps1)

	resp := protocol.Response{
		StepsMono:   steps1,
		StepsMulti:  stepsN,
		Speedup:     speedup,
		DurationSec: req.DurationSec,
		TopNodes:    topN,
	}

	encoder.Encode(resp)
}
