package main

import (
	"ELP/internal/protocol"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	workers := flag.Int("goroutines", 10, "nombre de goroutines (>= 1)")
	duration := flag.Int("duration", 30, "durée en secondes (>= 1)")
	graphName := flag.String("graph", "lyon", "graphe à utiliser (small|lyon)")

	flag.Parse()
	if *workers <= 0 {
		fmt.Println("Erreur : le nombre de goroutines doit être >= 1")
		os.Exit(1)
	}

	if *duration <= 0 {
		fmt.Println("Erreur : la durée doit être >= 1 seconde")
		os.Exit(1)
	}

	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)
	fmt.Println("Paramètres envoyés au serveur :")
	fmt.Println("Workers :", *workers)
	fmt.Println("Durée :", *duration, "secondes")
	if *graphName == "small" {
		fmt.Println("Graph utilisé: graph_small")
	} else {
		fmt.Println("Graph utilisé: graph_lyon")
	}
	req := protocol.Request{
		NumWalks:    *workers,
		DurationSec: *duration,
		Graph:       *graphName,
	}

	encoder.Encode(req)

	var resp protocol.Response
	decoder.Decode(&resp)

	fmt.Println("=== Résultat reçu du serveur pour la comparaison ===")
	fmt.Println("Durée :", resp.DurationSec, "secondes")
	fmt.Println("Nombre de pas total avec une goroutine :", resp.StepsMono)
	fmt.Println("Nombre de pas total avec N goroutines :", resp.StepsMulti)

	fmt.Printf("Rapport de pas : %.2f\n", resp.Speedup)
	fmt.Println("Top 5 nodes :")

	for i, n := range resp.TopNodes {
		fmt.Printf("%d) Node %d -> %.5f\n", i+1, n.Node, n.Prob)
	}
}
