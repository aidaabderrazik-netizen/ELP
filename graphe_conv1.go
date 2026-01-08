package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
)

func ChargerGraphe(nomFichier string) (map[int][]int, error) {
	file, err := os.Open(nomFichier)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	graphe := make(map[int][]int)
	reader := csv.NewReader(file)

	for {
		ligne, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(ligne) >= 2 {
			u, _ := strconv.Atoi(strings.TrimSpace(ligne[0]))

			voisinsRaw := strings.Trim(ligne[1], "\"")
			listeVoisins := strings.Split(voisinsRaw, ",")

			for _, vStr := range listeVoisins {
				vStr = strings.TrimSpace(vStr)
				if vStr == "" {
					continue
				}
				if v, errConv := strconv.Atoi(vStr); errConv == nil {
					graphe[u] = append(graphe[u], v)
				}
			}
		}
	}

	return graphe, nil
}

func fonc() {
	dico, _ := ChargerGraphe("lyon_graph.csv")

}
