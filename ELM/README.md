# Guess It! — Projet Elm
Fait par Malamba Ruth, Abderrazik Aida, Oaissa Adam
## Description

**Guess It** est un jeu de devinettes développé en **Elm**.  
Le joueur doit deviner un mot anglais à partir de ses définitions, récupérées dynamiquement via une API de dictionnaire.

Le jeu fonctionne avec un **timer**, un **score**, et un système de mots choisis aléatoirement.


##  API utilisée

Les définitions sont récupérées depuis : https://api.dictionaryapi.dev/api/v2/entries/en/{word}

Les données JSON sont décodées et transformées en structures Elm.



##  Fonctionnalités

-  Mot choisi aléatoirement depuis un fichier texte
-  Définitions regroupées par type grammatical
-  Timer de jeu
-  Score automatique
-  Saisie utilisateur
-  Affichage de la solution

##  Lancer le projet

1. Installer Elm  
   https://elm-lang.org/

2. Installer les dépendances :
```bash
elm make src/Main.elm
```
3. Lancer un serveur local :
```
elm reactor
```
4. Ouvrir :

http://localhost:8000
