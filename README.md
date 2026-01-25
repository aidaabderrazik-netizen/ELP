# Random Walk & PageRank 

Fait par Oaissa Adam, Abderrazick Aida, Malamba Ruth

## 1.Description du projet
Ce projet implémente une **marche aléatoire (Random Walk)** sur un graphe afin d’estimer une **distribution de probabilité de visite des noeuds**, principe utilisé dans l’algorithme **PageRank**.

Le projet met l’accent sur :
- la **concurrence en Go** (goroutines, channels, waitgroups),
- la **mesure de performance** entre **1 goroutine** et **N goroutines**,
- une **architecture client / serveur TCP**,
- l’utilisation de **graphes de grande taille** et de **graphes réduits** pour la visualisation.

Deux modes sont fournis :
1. **Mode test local** (sans réseau)
2. **Mode TCP client / serveur**

## 2.Principe algorithmique
#### Random Walk

Une marche aléatoire consiste à :

- partir d’un nœud initial

- choisir aléatoirement un voisin

- répéter le processus pendant une durée donnée

Chaque visite est comptabilisée.
La probabilité d’un nœud est estimée par :

probabilité(nœud) = nombre de visites du nœud / nombre total de visites


#### Mesure de performance
Pour comparer les performances, on utilise une durée fixe et on mesure :

- le **nombre total** de pas réalisés

- le **débit** (pas par seconde)

- le **rapport de performance** entre 1 goroutine et N goroutine

## 3. Version Test local

Cette version permet de comparer 1 goroutine vs N goroutines directement depuis la ligne de commande.

#### Lancement
Depuis la racine du projet :

go run ./test 

*(Par défaut, la durée est de 120 secondes et le nombres de goroutines pour comparaison est de 100. Pour modifier cela, il vous faut aller dans ./test/test.go ligne 43-44 )*
##### Resultats affichés
- temps écoulé

- nombre total de pas

- pas par seconde

- comparaison entre 1 goroutine et N goroutines

## 4.Version Client/Serveur
#### Principe
Le serveur écoute sur un port TCP et le client envoie :

- le nombre de goroutines

- la durée

- le graphe à utiliser



Le serveur limite le nombre de clients simultanés grâce à un sémaphore.

#### Lancement
Dans un premier terminal, lancez le serveur:

**go run ./TCP/serveur**

Dans un second, lancez le client :

**go run ./TCP/client -goroutines=8 -duration=30 -graph=small** 

Le nombre de connexion est limité à 4. 

#### Paramètres du client
- goroutines : nombre de goroutines (≥ 1) (50 par défaut)

- duration : durée en secondes (≥ 1) (30 sec par défaut )

- graph : graphe à utiliser parmis:

graph_small.csv (une dizaine de noeuds) et lyon_graph.csv (environ 3000 noeuds)

Par défaut, le gros graphe est utilisé, pour changer cela il faut ajouter -graph=small dans votre commande comme dans l'exemple. 



