// Plugin/HeartBeat.js
const uploader = require("./Uploader");
const vscode = require("vscode");

// Last heartbeat timestamp (ms), shared globally across all files
let lastHeartbeatTime = 0;

// ⏳ THE COOLDOWN (Standard is 2 minutes)
const COOLDOWN_MS = 120 * 1000;

/**
 * Handles the "Heartbeat" logic.
 * Decides if we should actually call the binary or just ignore the event.
 * @param {vscode.TextDocument} document
 */
function handleHeartbeat(document) {
  const now = Date.now();
  
  // 1. Check Cooldown
  if (now - lastHeartbeatTime < COOLDOWN_MS) return;

  // 2. Fire the Upload - returns if failed
  if (!uploader.spawnProcess(document)) return;

  // 3. Reset the Timer 
  lastHeartbeatTime = now;
}

module.exports = { handleHeartbeat };
