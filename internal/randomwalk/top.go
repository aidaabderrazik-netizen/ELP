package randomwalk

import (
	"sort"
)

// NodeProb représente un noeud du graphe associé
// à sa probabilité estimée de visite.
// Cette structure est utilisée pour le classement (Top K)
// et pour l’envoi des résultats au client (JSON).
type NodeProb struct {
	Node int64   `json:"node"` // Identifiant du noeud
	Prob float64 `json:"prob"` // Probabilité estimée de visite
}

// TopK retourne les k noeuds ayant les plus fortes probabilités.
// - probs : map[node] -> probabilité
// - k     : nombre de noeuds à retourner
func TopK(probs map[int64]float64, k int) []NodeProb {

	// Liste intermédiaire pour transformer la map
	// en slice (nécessaire pour pouvoir trier)
	var list []NodeProb

	// Conversion de la map en slice de NodeProb
	for n, p := range probs {
		list = append(list, NodeProb{n, p})
	}

	// Tri de la slice par probabilité décroissante
	sort.Slice(list, func(i, j int) bool {
		return list[i].Prob > list[j].Prob
	})

	// Si le nombre de noeuds est inférieur à k,
	// on retourne toute la liste
	if len(list) < k {
		return list
	}

	// Sinon, on retourne uniquement les k premiers
	// (les plus visités / les plus probables)
	return list[:k]
}
