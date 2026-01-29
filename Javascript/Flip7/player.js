// attribution joueur 


function randomInt(min, max) {
  return Math.floor(randomFloat(min, max + 1))
}

let liste_id =[]

function constructor(name) {
    // IL FAUT UNE LISTE DES IDENTIFIANTS DEJA EXISTANT
    let deck_player = []
    let points = 0 
    // let id = 1
    // while (id in liste_id) {
    //     id += 1
    // }
    let player = {
        "name" : name, 
        "deck_player" : deck_player,
        "points" : points,
        "manche" : true // initialement le joueur a le droit de jouer sa manche 
    }
    return player
}

console.log(constructor("jessica"))

function hasDuplicate(carte, deck) {
    for (let i=0; i<deck.lenght; i++ ) {
        if (carte == deck[i]) {
            return True // carte en doublon, il faut arreter la manche
        }
    }
}

function addCard(card, deck_player) {
    if (player[manche] == true) {
        carte = deck[randomInt(0, length(deck))]
        if (hasDuplicate(card, deck_player) == true) { 
            player[manche] = false
        }
    }
}

function resetHand(deck_player) {
    deck_player = []
}



