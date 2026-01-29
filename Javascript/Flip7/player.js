// attribution joueur 


function randomInt(min, max) {
  return Math.floor(randomFloat(min, max + 1))
}

let liste_id =[]

<<<<<<< HEAD
function constructor(name) {
=======
function createPlayer(id,name) {
>>>>>>> 9be65c1ea75f56d62766fdaccfcd625c4e7784f3
    // IL FAUT UNE LISTE DES IDENTIFIANTS DEJA EXISTANT
    let player = {
        "name" : name, 
<<<<<<< HEAD
        "deck_player" : deck_player,
        "points" : points,
=======
        "id" : id,
        "deck_player" : [],
        "totalPoints" : 0,
        "roundPoints": 0 ,
>>>>>>> 9be65c1ea75f56d62766fdaccfcd625c4e7784f3
        "manche" : true // initialement le joueur a le droit de jouer sa manche 
    }
    return player
}

<<<<<<< HEAD
console.log(constructor("jessica"))

function hasDuplicate(carte, deck) {
    for (let i=0; i<deck.lenght; i++ ) {
        if (carte == deck[i]) {
            return True // carte en doublon, il faut arreter la manche
=======
function hasDuplicate(carte, player_deck) {
    if (!isNaN(carte)) { // on verifie que c'est pas une carte bonus 
        if (player_deck.includes(carte)) {
                return true // carte en doublon, il faut arreter la manche
            }
>>>>>>> 9be65c1ea75f56d62766fdaccfcd625c4e7784f3
        }
    else {
        return false
    }

}

<<<<<<< HEAD
function addCard(card, deck_player) {
    if (player[manche] == true) {
        carte = deck[randomInt(0, length(deck))]
        if (hasDuplicate(card, deck_player) == true) { 
            player[manche] = false
        }
    }
}
=======
>>>>>>> 9be65c1ea75f56d62766fdaccfcd625c4e7784f3

function resetHand(deck_player) {
    deck_player = []
}

module.exports = {createPlayer ,resetHand,hasDuplicate}




