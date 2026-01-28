function createDeck() {
    const deck = []
    for (let valeur=1; valeur<=12; valeur++){
        for (let i=1; i<value; i++){
            deck.push({type:'number', value});
        }
    }
    const CARTES_ACTIONS = {SECOND:6, FREEZE:6, PLUS2:6, TIMES2:6};
    for (const [action,count] of Object.entries(CARTES_ACTIONS)) {
        for (let i=0; i<count; i++) {
            deck.push({type: 'action',action })
        }
    }
    return deck;
}
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
