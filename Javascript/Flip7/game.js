const { askQuestion, askPlayerChoice } = require("./input");
const { createPlayer, hasDuplicate } = require("./player");
const { drawCard } = require("./deck");
const { log } = require("./logger");

// ===================== START GAME =====================

async function askNumberOfPlayers() {
  let nb = 0;

  do {
    nb = parseInt(await askQuestion("Il y'a combien de joueurs (3 à 8) ? "), 10);
  } while (nb < 3 || isNaN(nb) || nb > 8);

  return nb;
}

async function createPlayers(nb) {
  let players = [];

  for (let id = 0; id < nb; id++) {
    const name = await askQuestion(`Nom du joueur n° ${id + 1}: `);
    players.push(createPlayer(id, name));
    log(`Joueur créé : ${name}`);
  }

  return players;
}

async function startGame() {
  console.log("========= Démarrage du jeu Flip7 =========");

  const nb_joueurs = await askNumberOfPlayers();
  console.log("Nombre de joueurs =", nb_joueurs);

  const players = await createPlayers(nb_joueurs);

  console.log("Le jeu commence !");
  return players;
}

// ===================== DISPLAY =====================

function showGameState(players) {
  console.log("\n================== État global après cette manche ==================");
  log("=== FIN DE LA MANCHE ===");

  players.forEach(player => {
    console.log(`- ${player.name} : Total = ${player.totalPoints} points`);
    log(`SCORE | ${player.name} : ${player.totalPoints}`);
  });

  console.log("=====================================================================\n");
}

// ===================== ROUND MANAGEMENT =====================

function stopRound(players) {
  players.forEach(p => p.manche = false);
}

function reinit(players) {
  players.forEach(p => {
    p.manche = true;
    p.deck_player = [];
    p.roundPoints = 0;
  });
}

function ifRound(players) {
  return players.some(p => p.manche);
}

// ===================== TURN DISPLAY =====================

function displayTurnHeader(player) {
  console.log(`\n---- La manche est à ${player.name} ----`);
  console.log(`Cartes actuelles : [${player.deck_player.join(",")}]`);
  console.log(`Score manche : ${player.roundPoints}`);
  console.log(`Score total : ${player.totalPoints}`);
}

// ===================== TURN ACTIONS =====================

function handleDuplicate(player, card, discardPile) {
  console.log(`Malheureusement c'est un doublon ${player.name}, vous gagnez 0 point.`);
  log(`${player.name} fait un doublon`);

  discardPile.push(...player.deck_player, card);

  player.manche = false;
  player.roundPoints = 0;
  player.deck_player = [];
}

function handleSevenCards(player, discardPile, players) {
  console.log(`${player.name} a tiré 7 cartes différentes ! Bonus +15 !`);
  log(`${player.name} atteint 7 cartes (+15 bonus)`);

  player.roundPoints += 15;
  player.totalPoints += player.roundPoints;

  discardPile.push(...player.deck_player);
  player.deck_player = [];

  stopRound(players);
  return 1;
}

function handleWin(player, players) {
  if (player.totalPoints >= 200) {
    console.log(`🎉 Félicitations ${player.name}, vous avez gagné avec ${player.totalPoints} points !`);
    log(`FIN DE PARTIE - VAINQUEUR : ${player.name}`);

    stopRound(players);
    return 1;
  }
  return 0;
}

function handleStop(player, discardPile) {
  console.log(`${player.name} s'arrête et garde ${player.roundPoints} points.`);
  log(`${player.name} stop avec ${player.roundPoints}`);

  player.totalPoints += player.roundPoints;

  discardPile.push(...player.deck_player);
  player.deck_player = [];

  player.manche = false;
}

// ===================== PLAY TURN =====================

async function handleDraw(player, deck, discardPile, players) {
  let card = drawCard(deck, discardPile);
  console.log(`${player.name} pioche : ${card}`);
  log(`${player.name} pioche ${card}`);

  if (hasDuplicate(card, player.deck_player)) {
    handleDuplicate(player, card, discardPile);
    return 0;
  }

  player.deck_player.push(card);
  player.roundPoints += card;

  if (player.deck_player.length === 7) {
    return handleSevenCards(player, discardPile, players);
  }

  return 0;
}

async function playTurn(player, deck, discardPile, players) {
  if (!player.manche) return 0;

  log(`Tour de ${player.name}`);
  displayTurnHeader(player);

  const cmd = await askPlayerChoice(player.name);

  if (cmd === "d") {
    const res = await handleDraw(player, deck, discardPile, players);

    // Vérifie si quelqu'un a gagné
    if (handleWin(player, players)) return 1;

    return res;
  }

  if (cmd === "s") {
    handleStop(player, discardPile);

    // Vérifie si quelqu'un a gagné après arrêt
    if (handleWin(player, players)) return 1;
  }

  return 0;
}

// ===================== END GAME =====================

function winGame(players) {
  for (let p of players) {
    if (p.totalPoints >= 200) {
      return [false, p.name];
    }
  }
  return [true, ""];
}

module.exports = {
  startGame,
  playTurn,
  ifRound,
  reinit,
  showGameState,
  winGame
};
