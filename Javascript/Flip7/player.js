constructor(name)

export function randomInt(min, max) {
  return Math.floor(randomFloat(min, max + 1))
}


export function hasDuplicate(carte, deck) {
    for (let i=0; i<deck.lenght; i++ ) {
        if (carte == deck[i]) {
            return True // carte en doublon, il faut arreter la manche
        }
    }
}

export function addCard(card, deck_player) {
    carte = deck[randomInt(0, length(deck))]
    if (hasDuplicate(card, deck_player) == True) { 
        print("La manche s'arrête") // manche arretée 
    }
}
    





resetHand()

// Module Joueur : fonction piocher, 

let deck_test = ["1", "2", "3"]




let carte = ""

