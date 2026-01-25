package main

import (
	"ELP/internal/protocol"   // Définit les structures Request / Response échangées avec le client
	"ELP/internal/randomwalk" // Contient la logique du random walk / PageRank
	"encoding/json"           // Sérialisation / désérialisation JSON
	"fmt"                     // Affichage console
	"net"                     // Réseau TCP
	"time"                    // Gestion des durées
)

// Canal utilisé comme sémaphore pour limiter le nombre de clients simultanés
// Ici : maximum 4 clients connectés en même temps
var maxClients = make(chan struct{}, 4)

// Graphe global (map noeud -> voisins)
var graph map[int64][]int64

func main() {

	// Création du serveur TCP en écoute sur le port 9000
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Serveur TCP en écoute sur le port 9000")

	// Boucle infinie : le serveur accepte les connexions entrantes
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// Chaque client est traité dans une goroutine séparée
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// Prend une place dans le sémaphore (bloque si 4 clients déjà connectés)
	maxClients <- struct{}{}
	// Libère la place à la fin du traitement
	defer func() { <-maxClients }()
	defer conn.Close()

	// Décodeur JSON pour lire la requête du client
	decoder := json.NewDecoder(conn)
	// Encodeur JSON pour envoyer la réponse
	encoder := json.NewEncoder(conn)

	var req protocol.Request
	// Lecture de la requête envoyée par le client
	if err := decoder.Decode(&req); err != nil {
		return
	}
	fmt.Println("Requête reçue :", req)

	// -------- Chargement du graphe --------
	var err error

	// Graphe par défaut
	filename := "lyon_graph.csv"

	// Si le client demande un petit graphe pour la visualisation
	if req.Graph == "small" {
		filename = "graph_small.csv"
	}

	// Chargement du graphe depuis le fichier CSV
	graph, err = randomwalk.ChargerGraphe(filename)
	if err != nil {
		panic(err)
	}

	// Injection du graphe dans le package randomwalk (variable globale interne)
	randomwalk.LoadGraph(graph)

	// S'assurer que tous les voisins existent comme clé dans la map
	// Évite les accès à des clés inexistantes pendant le random walk
	for _, voisins := range graph {
		for _, v := range voisins {
			if _, ok := graph[v]; !ok {
				graph[v] = []int64{}
			}
		}
	}

	// -------- Calcul des random walks --------

	// Conversion de la durée (en secondes) en time.Duration
	duration := time.Duration(req.DurationSec) * time.Second

	// Exécution avec 1 goroutine (référence mono-thread)
	steps1, _ := randomwalk.RunRandomWalks(1, duration)

	// Exécution avec N goroutines (paramétré par le client)
	stepsN, probsN := randomwalk.RunRandomWalks(req.NumWalks, duration)

	// Extraction des K noeuds les plus visités (ici top 5)
	topN := randomwalk.TopK(probsN, 5)

	// Calcul du facteur d'accélération (option B : nombre de pas)
	speedup := float64(stepsN) / float64(steps1)

	// Construction de la réponse envoyée au client
	resp := protocol.Response{
		StepsMono:   steps1,          // Nombre total de pas avec 1 goroutine
		StepsMulti:  stepsN,          // Nombre total de pas avec N goroutines
		Speedup:     speedup,         // Rapport stepsN / steps1
		DurationSec: req.DurationSec, // Durée du test
		TopNodes:    topN,            // Top 5 des noeuds les plus visités
	}

	// Envoi de la réponse au client au format JSON
	encoder.Encode(resp)
}
