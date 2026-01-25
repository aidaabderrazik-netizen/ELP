package main

import (
	"ELP/internal/protocol" // Structures Request / Response partagées avec le serveur
	"encoding/json"         // Encodage / décodage JSON pour la communication réseau
	"flag"                  // Gestion des paramètres passés en ligne de commande
	"fmt"                   // Affichage console
	"net"                   // Connexion TCP
	"os"                    // Gestion des sorties du programme (Exit)
)

func main() {
	// Définition des paramètres en ligne de commande
	// -goroutines : nombre de goroutines à utiliser côté serveur
	// -duration   : durée du test en secondes
	// -graph      : type de graphe à analyser (small ou lyon)
	workers := flag.Int("goroutines", 50, "nombre de goroutines (>= 1)")
	duration := flag.Int("duration", 30, "durée en secondes (>= 1)")
	graphName := flag.String("graph", "lyon", "graphe à utiliser (small|lyon)")

	// Analyse des paramètres fournis par l'utilisateur
	flag.Parse()

	// Vérification de la validité du nombre de goroutines
	if *workers <= 0 {
		fmt.Println("Erreur : le nombre de goroutines doit être >= 1")
		os.Exit(1)
	}

	// Vérification de la validité de la durée
	if *duration <= 0 {
		fmt.Println("Erreur : la durée doit être >= 1 seconde")
		os.Exit(1)
	}

	// Connexion au serveur TCP local sur le port 9000
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Création de l’encodeur et du décodeur JSON sur la connexion TCP
	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	// Affichage des paramètres envoyés au serveur
	fmt.Println("Paramètres envoyés au serveur :")
	fmt.Println("Workers :", *workers)
	fmt.Println("Durée :", *duration, "secondes")

	if *graphName == "small" {
		fmt.Println("Graph utilisé: graph_small")
	} else {
		fmt.Println("Graph utilisé: graph_lyon")
	}

	// Construction de la requête à envoyer au serveur
	req := protocol.Request{
		NumWalks:    *workers,   // Nombre de goroutines
		DurationSec: *duration,  // Durée du test
		Graph:       *graphName, // Type de graphe
	}

	// Envoi de la requête au serveur
	encoder.Encode(req)

	// Réception de la réponse du serveur
	var resp protocol.Response
	decoder.Decode(&resp)

	// Affichage des résultats reçus
	fmt.Println("=== Résultat reçu du serveur pour la comparaison ===")
	fmt.Println("Durée :", resp.DurationSec, "secondes")
	fmt.Println("Nombre de pas total avec une goroutine :", resp.StepsMono)
	fmt.Println("Nombre de pas total avec N goroutines :", resp.StepsMulti)

	// Affichage du rapport de performance (option B)
	fmt.Printf("Rapport de pas : %.2f\n", resp.Speedup)
	fmt.Println("Top 5 nodes :")

	// Affichage du top 5 des nœuds les plus visités
	for i, n := range resp.TopNodes {
		fmt.Printf("%d) Node %d -> %.5f\n", i+1, n.Node, n.Prob)
	}
}
