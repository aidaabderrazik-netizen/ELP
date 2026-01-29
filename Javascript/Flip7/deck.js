
function createDeck() {
    const deck = [0]
    for (let value=1; value<=12; value++){
        for (let i=1; i<=value; i++){
            deck.push(value);
        }
    }
    const CARTES_ACTIONS = ["SECOND", "FREEZE", "PLUS2", "TIMES2"]
    for (let i=0; i<=3; i++){
        for (let j=1; j<=6; j++){
            deck.push(CARTES_ACTIONS[i]);
        }
    }

    return deck;
}

// Algorithme de Fisher-Yates 

function shuffle(deck) {
    for (let i=deck.length-1; i>0; i--) {
        const j= Math.floor(Math.random()*(i+1));
        [deck[i],deck[j]]=[deck[j], deck[i]];
    }
    return deck;
}

function drawCard(deck) {
    if (deck.length === 0) {
        throw new Error('Le paquet est vide');
    }
    return deck.pop();

}
