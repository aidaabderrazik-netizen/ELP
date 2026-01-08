package main

import (
	"encoding/csv"
	"fmt"
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
			listeDesVoisins := strings.Split(voisinsRaw, ",")

			for _, vStr := range listeDesVoisins {
				vStr = strings.TrimSpace(vStr)
				if vStr == "" {
					continue
				}
				v, errConv := strconv.Atoi(vStr)
				if errConv == nil {
					graphe[u] = append(graphe[u], v)
				}
			}
		}
	}

	return graphe, nil 
}

func main() {
	monGraphe, err := ChargerGraphe("data.csv")
	if err != nil {
		fmt.Println("Erreur lors du chargement :", err)
		return
	}
	fmt.Printf("Le sommet 2 possède %d voisins.\n", len(monGraphe[2]))
}
