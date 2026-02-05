
function createDeck() { 
    const deck = [0]
    for (let value=1; value<=12; value++){
        for (let i=1; i<=value; i++){
            deck.push(value);
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

function drawCard(deck,discardPile) {
    if (deck.length === 0) {
        console.log('Le paquet est vide, on mélange la défausse ...');
        deck.push(...shuffle(discardPile));
        discardPile.length = 0;

    }
    return deck.pop();

}


module.exports = {createDeck,shuffle,drawCard}
