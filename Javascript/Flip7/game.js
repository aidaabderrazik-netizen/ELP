
const { askQuestion, askPlayerChoice } = require("./input")
const { createPlayer, hasDuplicate } = require("./player")
const { drawCard } = require("./deck")
const { initLogger, log } = require("./logger");

async function startGame(){
    let players = [];
    //combien de joueurs 
    console.log("=========Démarrage du jeu Flip7=========")
    let nb_joueurs= 0;
    do {
        nb_joueurs= parseInt(await askQuestion("Il y'a combien de joueurs (compris entre 3 et 8 inclus ) ?"),10);
    } while (nb_joueurs < 3 || isNaN(nb_joueurs) || nb_joueurs > 8  )
    
    console.log("Nombre de joueurs=", nb_joueurs);
    for (let id=0 ; id<nb_joueurs; id++) {
        const player_name = await askQuestion(`Nom du joueur n° ${id+1}:`);
        players.push(createPlayer(id,player_name));
        log(`Joueur crée : ${player_name}`)


    }

    console.log("Le jeu commence ");
    //console.log(players)
    return players
}

//affiche l'état du jeu 
function showGameState(players) {
    console.log("\n======================== État global après cette manche ============================");
    log(`=== FIN DE LA MANCHE ===`);
    players.forEach(player => {
        console.log(`- ${player.name} : Total = ${player.totalPoints} points. `);
        log(`SCORE | ${player.name} : ${player.totalPoints} points`);

    });



    console.log("=====================================================================================\n");
}


// pour un joueur pendant une manche
async function playTurn(player, deck, discardPile,players) {
    
    if (player.manche) {
        log(`Tour de ${player.name}`);
        console.log(`\n---- La manche est à ${player.name}---`);
        console.log(`Cartes actuelles : [${player.deck_player.join(",")}]`);
        console.log(`Score actuel de la manche: ${player.roundPoints}`);
        console.log(`Score total (sans la manche) : ${player.totalPoints}`);
        const cmd = await askPlayerChoice(player.name); // demande le choix 
        

        if (cmd==="d") {
            let card = drawCard(deck, discardPile);
            log(`${player.name} pioche ${card}`);
            console.log(`${player.name} pioche : ${card}`);

            if ( hasDuplicate(card, player.deck_player) == true) {
                console.log(`Malheureusement c'est un doublon, ${player.name}, vous arretez la manche et vous gagnez 0 points.`) 
                player.manche=false ;
                player.roundPoints= 0 ;
                discardPile.push(...player.deck_player,card);
                player.deck_player= [];
                log(`${player.name} fait un doublon et perd sa manche`);
                
                return 0
            }
            else {
                player.deck_player.push(card);
                player.roundPoints += card ;
                
                
                 // verifie si le joueur a tiré 7 cartes
                if (player.deck_player.length ===7) {
                    console.log(`${player.name} a tiré 7 cartes différentes et obtient un bonus de 15 points. Cette manche est finie`);
                    player.roundPoints += 15; // ajoute le bonus dans roundPoints
                    player.totalPoints += player.roundPoints; // transfert final
                    discardPile.push(...player.deck_player);
                    player.deck_player = [];
                    log(`${player.name} atteint 7 cartes et gagne un bonus de 15 points`);
                    player.manche=false;
                    stopRound(players);
                    return 1;
                }
                // si le joueur atteint le score gagnant, on transfère ses points de manche au total
                if (player.totalPoints + player.roundPoints >= 100) {
                    console.log(`Félicitations ${player.name}, vous cummulez  au moins 100 points!`)
                    log(`${player.name} atteint ${player.totalPoints} points et déclenche la fin de partie`);
                    player.totalPoints += player.roundPoints;
                    player.manche = false;
                    stopRound(players);
                    return  1 ;
            }
                
            }
        }
        if (cmd === "s") {
            console.log(`${player.name} s'arrête et garde ${player.roundPoints} points`);
            player.manche=false ;
            discardPile.push(...player.deck_player);
            player.deck_player = [];
            player.totalPoints += player.roundPoints; 
            log(`${player.name} s'arrête avec ${player.roundPoints} points pour cette manche, et un score total de ${player.totalPoints}`);
        }
    } 
    return 0
}

function stopRound(players) {
    players.forEach(p => p.manche = false);
}

async function reinit(players){
    //Pour la réinitialisation des manches
    for (let i=0 ; i<players.length;i++){
        players[i].manche=true;
        players[i].deck_player=[];
        players[i].roundPoints=0;

    }
}

//détermine si le jeu est finie (à 100 points par convention)
function winGame(players){
    for ( let i=0; i<players.length; i++ ) {
        if (players[i].totalPoints >= 100) {
            log(`FIN DE PARTIE - VAINQUEUR : ${players[i].name}`);
            return [false,players[i].name];
        }
    }
    return [true,""]

}

//verifie que la manche continue
function ifRound(players) {
    for ( let i=0; i<players.length; i++ )  {
        if (players[i].manche){ // au moins un joueur est encore dans la manche
            return true ;
        }
    }
    reinit(players)
    return false
}


module.exports = {startGame, playTurn, ifRound, reinit , showGameState, winGame}
