// attribution joueur 


function randomInt(min, max) {
  return Math.floor(randomFloat(min, max + 1))
}

let liste_id =[]

function createPlayer(id,name) {
    // IL FAUT UNE LISTE DES IDENTIFIANTS DEJA EXISTANT
    let player = {
        "name" : name, 
        "id" : id,
        "deck_player" : [],
        "totalPoints" : 0,
        "roundPoints": 0 ,
        "manche" : true // initialement le joueur a le droit de jouer sa manche 
    }
    return player
}

function hasDuplicate(carte, player_deck) {
    if (!isNaN(carte)) { // on verifie que c'est pas une carte bonus 
        if (player_deck.includes(carte)) {
                return true // carte en doublon, il faut arreter la manche
            }
        }
    else {
        return false
    }

}


function resetHand(deck_player) {
    deck_player = []
}

module.exports = {createPlayer ,resetHand,hasDuplicate}




