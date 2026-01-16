// Plugin/Setup.js
const vscode = require("vscode");
const env = require("./Config");
const downloader = require("./BinaryDownload");
const statusHelper = require("./StatusBarUpdate"); // Renamed file
const fs = require("fs");
const path = require("path");
const os = require("os");

async function runSetup(statusBar) {
  const config = env.getConfig() || {};
  const currentUri = config.MONGO_URI || "";

  const uri = await vscode.window.showInputBox({
    placeHolder: "mongodb+srv://admin:password@...",
    prompt: "Enter (or update) your MongoDB Connection String",
    value: currentUri,
    ignoreFocusOut: true,
    password: true,
  });

  if (uri === undefined) return;

  const homeDir = os.homedir();
  const configPath = path.join(homeDir, ".takatime.json");

  let newConfig = {
    MONGO_URI: uri,
    VERSION: config.VERSION || env.CURRENT_VERSION,
  };

  try {
    fs.writeFileSync(configPath, JSON.stringify(newConfig, null, 4));
    vscode.window.showInformationMessage("Configuration Saved! ✅");
  } catch (e) {
    vscode.window.showErrorMessage("Failed to save config");
    return;
  }

  // Check & Download Binary
  const isBinaryReady = env.checkBinary(newConfig.VERSION);
  if (!isBinaryReady) {
    try {
      const success = await downloader.downloadBinary(newConfig.VERSION);
      if (success) {
        vscode.window.showInformationMessage(
          `TakaTime ${newConfig.VERSION} installed successfully! 🚀`
        );
      }
    } catch (err) {
      vscode.window.showErrorMessage(`Download Failed: ${err.message}`);
      return;
    }
  }

  // Update Status Bar
  statusHelper.checkStatus(statusBar);
}

module.exports = { runSetup };
