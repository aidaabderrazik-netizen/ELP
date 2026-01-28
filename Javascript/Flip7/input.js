const readline = require("readline")

const rl =readline.createInterface({input: process.stdin, output :process.stdout}) //permet de recuperer la demande de l'utilisateur


// fonction qui pinermet d'afficher une question
function askQuestion(question) {
  return new Promise((resolve) => {
    rl.question(question, (answer) => {
      resolve(answer.trim());
    });
  });
}



function closeInput() {
  rl.close();
}

module.exports = { askQuestion, closeInput };
//retourne "draw" ou "stop"