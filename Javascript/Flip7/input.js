const { createInterface } = require("readline");

const rl =createInterface({input: process.stdin, output :process.stdout}) //permet de recuperer la demande de l'utilisateur


// fonction qui pinermet d'afficher une question
function askQuestion(question) {
  return new Promise((resolve) => {
    rl.question(question, (answer) => {
      resolve(answer.trim());
    });
  });
}



// Savoir les intentions d'un joueur 
async function askPlayerChoice(playerName){
    let choice ;
    do {
        choice = await askQuestion(`${playerName}, tirer (d) ou s'arrêter (s) ?`);
    } while (choice !== "d" && choice !=="s")
    return choice ;
    }




function closeInput() {
  rl.close();
}

module.exports = { askQuestion, closeInput, askPlayerChoice};
