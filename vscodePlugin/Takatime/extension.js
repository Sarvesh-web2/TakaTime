const vscode = require("vscode");
const path = require("path");
const statusHelper = require("./Plugin/StatusBarUpdate");
const setupHelper = require("./Plugin/Setup");
// 👇 CHANGE THIS IMPORT
const heartbeat = require("./Plugin/HeartBeat");

/**
 * @param {vscode.ExtensionContext} context
 */
async function activate(context) {
  console.log("TakaTime: Initializing...");

  // 1. Status Bar
  const statusBar = vscode.window.createStatusBarItem(
    vscode.StatusBarAlignment.Left,
    100
  );
  statusBar.text = "$(sync~spin) TakaTime: Checking...";
  statusBar.command = "takatime.setup";
  statusBar.show();
  context.subscriptions.push(statusBar);

  // 2. Setup Command
  const setupCommand = vscode.commands.registerCommand("takatime.setup", () => {
    setupHelper.runSetup(statusBar);
  });
  context.subscriptions.push(setupCommand);

  // 3. ⚡ SAVE LISTENER (Now with Heartbeat Logic!)
  const saveListener = vscode.workspace.onDidSaveTextDocument((document) => {
    // Filter out junk
    if (document.uri.scheme !== "file") return;
    if (document.fileName.includes(path.sep + ".git" + path.sep)) return;

    // 👇 CALL THE HEARTBEAT MANAGER
    heartbeat.handleHeartbeat(document);
  });

  context.subscriptions.push(saveListener);

  // 4. Initial Check
  statusHelper.checkStatus(statusBar);
}

function deactivate() {}

module.exports = { activate, deactivate };
