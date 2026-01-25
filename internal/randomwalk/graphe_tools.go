package randomwalk

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
)

// graph contient la structure du graphe en mémoire
// clé   : id du noeud
// valeur: liste des voisins accessibles depuis ce noeud
var graph map[int64][]int64

// nodes contient la liste de tous les noeuds du graphe
// utile pour tirer un noeud aléatoire
var nodes []int64

// ChargerGraphe lit un fichier CSV et construit le graphe
// Format attendu du csv :
// node,neighbors
// 143403,"21714981,21718288,143408"
func ChargerGraphe(nomFichier string) (map[int64][]int64, error) {

	// Ouverture du fichier CSV
	file, err := os.Open(nomFichier)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Création de la map du graphe
	graphe := make(map[int64][]int64)

	// Lecteur CSV
	reader := csv.NewReader(file)

	for {
		// Lecture ligne par ligne
		ligne, err := reader.Read()
		if err == io.EOF {
			break // fin du fichier
		}
		if err != nil {
			return nil, err
		}

		// Ignorer l'en-tête du CSV
		if ligne[0] == "node" {
			continue
		}

		// Vérifie qu'on a au moins node + neighbors
		if len(ligne) >= 2 {

			// Conversion de l'id du noeud
			u, errU := strconv.ParseInt(strings.TrimSpace(ligne[0]), 10, 64)
			if errU != nil {
				continue // ligne invalide
			}

			// Suppression des guillemets autour de la liste de voisins
			voisinsRaw := strings.Trim(ligne[1], "\"")

			// Séparation des voisins
			listeVoisins := strings.Split(voisinsRaw, ",")

			// Conversion de chaque voisin
			for _, vStr := range listeVoisins {
				vStr = strings.TrimSpace(vStr)
				if vStr == "" {
					continue
				}

				// Ajout du voisin si la conversion réussit
				if v, errConv := strconv.ParseInt(vStr, 10, 64); errConv == nil {
					graphe[u] = append(graphe[u], v)
				}
			}
		}
	}

	// Retourne le graphe construit
	return graphe, nil
}

// LoadGraph charge le graphe dans le package randomwalk
// et initialise la liste des noeuds
func LoadGraph(g map[int64][]int64) {
	graph = g
	initNodes()
}

// initNodes construit la slice "nodes"
// contenant tous les noeuds du graphe
func initNodes() {
	nodes = make([]int64, 0, len(graph))
	for node := range graph {
		nodes = append(nodes, node)
	}
}
