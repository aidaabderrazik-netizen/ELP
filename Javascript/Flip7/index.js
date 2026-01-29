<<<<<<< HEAD
// sert à lancer le jeu → main
=======
// sert à lancer le jeu
const { startGame, playTurn , ifRound , showGameState, winGame, reinit} = require("./game");
const { closeInput } = require("./input");
const { createPlayer } = require("./player");
const { createDeck, shuffle } = require("./deck");
const { initLogger, log } = require("./logger");

async function main() {
  initLogger(); //écriture
  log("Démarrage de la partie");

  let deck = createDeck();
  deck= shuffle(deck)

  //console.log(deck);
  let players = await startGame();
  let run = true ;
  let winner = "";
  let discardPile = []; //la défausse
  let startIndex=0 ; // sert à gérer la rotation des joueurs
  do {
    console.log(`==============================NOUVELLE MANCHE========================`)
    reinit(players)
    log(`----------Nouvelle manche----------`)
    let res = 0; // 1 si un joueur a pioché 7 cartes
    // la manche démarre 
    do {

      for (let offset= 0 ; offset < players.length; offset++) {
        const i = (startIndex + offset) % players.length ;
        const res = await playTurn(players[i],deck,discardPile,players);
        if (res===1) break;
      }
    }   while (ifRound(players) && res ===0 )
    // fin de la manche
    showGameState(players)
    console.log(`Deck: ${deck.length} cartes | Défausse: ${discardPile.length} cartes`);

    

    const verif =winGame(players) //si un joueur a plus de 100 points run = false
    run=verif[0]
    winner=verif[1]
  } while (run)
  
    console.log(`\n=== Jeu terminé ! Le/La gagnant(e) est ${winner} ! ======`);

  
  closeInput();
}



main();
>>>>>>> 9be65c1ea75f56d62766fdaccfcd625c4e7784f3
