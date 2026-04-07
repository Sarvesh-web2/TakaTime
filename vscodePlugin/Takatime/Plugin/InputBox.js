const vscode = require("vscode");
const fs = require("fs");
const path = require("path");
const os = require("os");

const openInpuToGetMongoUri = async () => {
  try {
    const configPath = path.join(os.homedir(), ".takatime.json");
    // Ask user for URI
    const uri = await vscode.window.showInputBox({
      placeHolder: "mongodb+srv://admin:password@...",
      prompt: "Enter your MongoDB Connection String to start tracking",
      ignoreFocusOut: true,
      password: true, // 🔒 Hides the text for privacy (optional)
    });

    if (!uri) {
      vscode.window.showWarningMessage(
        "TakaTime: MONGO_URI is required to proceed.",
      );
      return null;
    }

    if (!uri.startsWith("mongodb")) {
      vscode.window.showErrorMessage("TakaTime: Invalid MongoDB URI format.");
      return null;
    }
    //regex to verfify local and hosted mongo uris
    const mongoUriRegex = /^(mongodb(?:\+srv)?):\/\/.+/;
    if (!mongoUriRegex.test(uri)) {
      vscode.window.showErrorMessage("TakaTime: Invalid MongoDB URI format.");
      return null;
    }
    addUriToConfig(uri, configPath);
    vscode.window.showInformationMessage(
      "TakaTime: MONGO_URI saved. Please restart VSCode to continue.",
    );
  } catch (err) {
    vscode.window.showErrorMessage(
      `TakaTime: Error getting MONGO_URI: ${err.message}`,
    );
    return null;
  }
};

const addUriToConfig = (uri, configPath) => {
  try {
    const rawConfig = fs.readFileSync(configPath, "utf8");
    const config = JSON.parse(rawConfig);
    config.MONGO_URI = uri;
    fs.writeFileSync(configPath, JSON.stringify(config, null, 4));
    vscode.window.showInformationMessage(
      "TakaTime: MONGO_URI saved to config.",
    );
  } catch (err) {
    vscode.window.showErrorMessage(
      `TakaTime: Failed to update config: ${err.message}`,
    );
  }
};
