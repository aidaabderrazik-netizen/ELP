// qui joue ? quand 
const { askQuestion } = require("./input")
const { constructor } = require("./player")

async function startGame(){
    let players = [];
    //combiien de joueurs 
    console.log("=========Démarrage du jeu Flip7=========")
    const nb_joueurs= parseInt(await askQuestion("Il y'a combien de joueurs ?"),10);
    console.log("Nombre de joueurs=", nb_joueurs);
    for (let id=0 ; id<nb_joueurs; id++) {
        const player_name = await askQuestion(`Nom du joueur n° ${id+1}:`);
        players.push(constructor(player_name));
        

    }
    console.log("Le jeu commence ")
}


//playTurn(player)


//endGame()


//isGameOver()

module.exports = { startGame}
