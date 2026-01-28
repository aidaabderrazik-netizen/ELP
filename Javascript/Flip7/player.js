constructor(name)


addCard(card)


hasDuplicate()


resetHand()

// Module Joueur : fonction piocher, 

let deck_test = ["1", "2", "3"]

export function randomInt(min, max) {
  return Math.floor(randomFloat(min, max + 1))
}

export function comparaison(carte, deck) {
    for (let i=0; i<deck.lenght; i++ ) {
        if (carte == deck[i]) {
            return True // carte en doublon, il faut arreter la manche
        }
    }
}

let carte = ""

export function tirer_carte(deck, deck_joueur) {
    carte = deck[randomInt(0, length(deck))]

    if (comparaison(carte, deck_joueur) == True) { 
        print("La manche s'arrête") // manche arretée 
    }
}
    
