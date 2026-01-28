const fs = require('fs');
const path = require('path');


// Chemin du dossier et du fichier de log
const LOG_DIR = path.join(__dirname, 'logs');
const LOG_FILE = path.join(LOG_DIR, 'game_log.txt');


function log(message) {
  const line = message + '\n';
  fs.appendFileSync(LOG_FILE, line, { encoding: 'utf8' });
}
