const fs = require('fs');
const path = require('path');


// Chemin du dossier et du fichier de log
const LOG_DIR = path.join(__dirname, 'logs');
const LOG_FILE = path.join(LOG_DIR, 'game_log.txt');


// initialise le fichier au lancement du jeu
function initLogger() {
    fs.writeFileSync(LOG_FILE, "===== NOUVELLE PARTIE FLIP7 =====\n", { flag: "w" });
}

function log(message) {
  const line = message + '\n';
  fs.appendFileSync(LOG_FILE, line, { encoding: 'utf8' });
}

module.exports={log,initLogger}
