# Flip7 — Jeu de cartes en JavaScript (Node.js)

Fait par Oaissa Adam, Malamba Ruth, Abderrazik Aida 

## Principe 
Flip7 est un jeu de cartes en ligne de commande développé en JavaScript (Node.js).
Le but du jeu est d’atteindre **100 points ou plus** en accumulant des points sur plusieurs manches, tout en évitant les doublons.

---

##  Prérequis

- **Node.js** (version 16 ou supérieure recommandée)

Vérification :
```bash
node -v
```
## Lancer jeu
Le jeu se lance en ligne de commande et guide les joueurs pas à pas. 
``` bash
npm start
```

## Règles du jeu
- Entre 3 et 8 joueurs
- Chaque manche :
    - Un joueur peut :
       - d → piocher une carte
        - s → s’arrêter et conserver ses points de manche
- Si un joueur :
    - Pioche un **doublon** → il perd tous les points de la manche
    - Pioche **7 cartes différentes** → la manche s’arrête et il gagne **+15 points**
- Les cartes jouées sont envoyées dans une défausse
- Le jeu s’arrête dès qu’un joueur atteint 100 points ou plus

## Fichier d'enregistrement des manches (game_log)
Toutes les actions importantes du jeu sont enregistrées dans le fichier : game_log.txt.

Chaque ligne correspond à un événement du jeu.

### Evenements enregistrés

- Création des joueurs
- Tours de jeu
- Pioches de cartes
- Doublons
- Arrêts volontaires
- Fin de manche
- Scores
- Victoire finale

### Fin de partie
La partie se termine automatiquement lorsqu’un joueur atteint 100 points ou plus.
Le vainqueur est affiché à l’écran et enregistré dans le fichier de log.

### Remarques
Le module disponible de Flip7 n'a pas été utilisé pour réaliser ce projet. 
L'IA a été utilisé pour le deboggage. 
La gestion des cartes spéciales n'est pas incluse dans notre module. 

