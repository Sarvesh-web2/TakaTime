// Plugin/Heartbeat.js
const uploader = require("./Uploader");
const vscode = require("vscode");

// Key: File Path, Value: Last Timestamp (ms)
let lastHeartbeatTime = 0;

// ⏳ THE COOLDOWN (Standard is 2 minutes)
const COOLDOWN_MS = 120 * 1000;

/**
 * Handles the "Heartbeat" logic.
 * Decides if we should actually call the binary or just ignore the event.
 * @param {vscode.TextDocument} document
 */
function handleHeartbeat(document) {
  const filePath = document.fileName;
  const now = Date.now();
  
  // 1. Check Cooldown
  if (now - lastHeartbeatTime < COOLDOWN_MS) return;

  // 2. Fire the Upload
  uploader.spawnProcess(document);

  // 3. Reset the Timer
  lastHeartbeatTime = now;
}

module.exports = { handleHeartbeat };
