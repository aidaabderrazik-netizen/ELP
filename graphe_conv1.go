package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
)

func ChargerGraphe(nomFichier string) (map[int64][]int64, error) {
	file, err := os.Open(nomFichier)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	graphe := make(map[int64][]int64)
	reader := csv.NewReader(file)

	for {
		ligne, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Ignorer le header
		if ligne[0] == "node" {
			continue
		}

		if len(ligne) >= 2 {
			u, errU := strconv.ParseInt(strings.TrimSpace(ligne[0]), 10, 64)
			if errU != nil {
				continue
			}

			voisinsRaw := strings.Trim(ligne[1], "\"")
			listeVoisins := strings.Split(voisinsRaw, ",")

			for _, vStr := range listeVoisins {
				vStr = strings.TrimSpace(vStr)
				if vStr == "" {
					continue
				}
				if v, errConv := strconv.ParseInt(vStr, 10, 64); errConv == nil {
					graphe[u] = append(graphe[u], v)
				}
			}
		}
	}

	return graphe, nil
}

// fonction qui retourne le graphe chargé depuis le CSV
func fonc(nomFichier string) (map[int64][]int64, error) {
	return ChargerGraphe(nomFichier)
}
