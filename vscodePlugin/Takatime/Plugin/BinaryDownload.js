const fs = require("fs");
const path = require("path");
const os = require("os");
const https = require("https");
const vscode = require("vscode");

// 1. Updated to accept the base name (e.g., "taka-upload" or "taka-dashboard")
function getPlatformFilename(baseName) {
  const plat = process.platform;
  const arch = process.arch;

  let osStr = "";
  if (plat === "win32") osStr = "windows";
  else if (plat === "linux") osStr = "linux";
  else if (plat === "darwin") osStr = "darwin";
  else return null;

  let archStr = "";
  if (arch === "x64") archStr = "amd64";
  else if (arch === "arm64") archStr = "arm64";
  else return null;

  const ext = plat === "win32" ? ".exe" : "";

  // OUTPUT EXAMPLE: taka-upload-windows-amd64.exe
  return `${baseName}-${osStr}-${archStr}${ext}`;
}

// 2. Extracted the actual HTTP download logic into a clean, reusable Promise
function downloadSingleBinary(url, destPath) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(destPath);

    const request = (uri) => {
      https
        .get(uri, (response) => {
          // Handle Redirects
          if (response.statusCode === 301 || response.statusCode === 302) {
            return request(response.headers.location);
          }

          // Handle Errors (like 404s)
          if (response.statusCode !== 200) {
            reject(new Error(`HTTP ${response.statusCode}`));
            return;
          }

          response.pipe(file);

          file.on("finish", () => {
            file.close(() => {
              // Make executable on Linux/Mac
              if (process.platform !== "win32") {
                try {
                  fs.chmodSync(destPath, 0o755);
                } catch (e) {
                  console.error("Chmod failed", e);
                }
              }
              resolve(true);
            });
          });
        })
        .on("error", (err) => {
          fs.unlink(destPath, () => {}); // Delete partial file on error
          reject(err);
        });
    };

    request(url);
  });
}

// 3. The main function that loops through all required binaries
async function ensureBinaries(version) {
  // List all the binaries you want to download here
  const requiredBinaries = ["taka-upload", "taka-dashboard"];

  const homeDir = os.homedir();
  const binDir = path.join(homeDir, ".takatime", "bin");

  // Ensure directory exists
  if (!fs.existsSync(binDir)) {
    fs.mkdirSync(binDir, { recursive: true });
  }

  // Wrap the whole process in one smooth VS Code loading notification
  return vscode.window.withProgress(
    {
      location: vscode.ProgressLocation.Notification,
      title: `Installing TakaTime ${version}...`,
      cancellable: false,
    },
    async (progress) => {
      try {
        for (const baseName of requiredBinaries) {
          const filename = getPlatformFilename(baseName);
          if (!filename) {
            throw new Error(
              `Unsupported Platform (${process.platform}-${process.arch})`,
            );
          }

          const ext = process.platform === "win32" ? ".exe" : "";
          // Saves as: taka-upload-v2.2.0.exe AND taka-dashboard-v2.2.0.exe
          const localFilename = `${baseName}-${version}${ext}`;
          const destPath = path.join(binDir, localFilename);
          const url = `https://github.com/Rtarun3606k/TakaTime/releases/download/${version}/${filename}`;

          // Update the UI text so the user knows which file is currently downloading
          progress.report({ message: `Fetching ${baseName}...` });

          // If a broken/old file is sitting there, delete it first
          if (fs.existsSync(destPath)) {
            fs.unlinkSync(destPath);
          }

          // Await the download before moving to the next binary
          await downloadSingleBinary(url, destPath);
        }

        vscode.window.showInformationMessage(
          `TakaTime binaries updated successfully!`,
        );
        return true;
      } catch (error) {
        vscode.window.showErrorMessage(
          `TakaTime Update Failed: ${error.message}`,
        );
        return false;
      }
    },
  );
}

// Export the newly named function
module.exports = { ensureBinaries };
