# ⏳ TakaTime (VS Code Edition)

**A privacy-first, serverless time tracker for your GitHub Profile.**

TakaTime is a lightweight alternative to WakaTime that requires **no monthly fees** and **no external servers**. It tracks your coding activity directly from VS Code and saves it to your own free MongoDB Atlas cluster.

![Dashboard Preview](https://github.com/Rtarun3606k/TakaTime/blob/main/assets/dashboard-preview.png?raw=true)

## 🚀 Features

- **Serverless:** No hosting required. Runs entirely via VS Code and GitHub Actions.
- **Free Forever:** Uses the MongoDB Atlas free tier (512MB storage = years of data).
- **Privacy-First:** You own your data. No third-party analytics.
- **Auto-Sync:** Updates your GitHub Profile README with your latest stats.
- **Language Tracking:** Automatically detects languages (Python, Go, JS, etc.) and projects.

## 📦 Installation

1.  Install this extension from the VS Code Marketplace.
2.  Open VS Code. The extension will automatically check for the necessary binary.
3.  If this is your first time, you will be prompted to enter your **MongoDB Connection String**.

## ⚙️ Setup Guide

### Step 1: Get a Database

1.  Create a free account on [MongoDB Atlas](https://www.mongodb.com/atlas/database).
2.  Create a new Cluster (Free Tier).
3.  Go to **Database Access** -> Add a new user (keep the password safe!).
4.  Go to **Network Access** -> Allow access from anywhere (`0.0.0.0/0`).
5.  Click **Connect** -> **Drivers** -> Copy the connection string (e.g., `mongodb+srv://user:pass@cluster0...`).

### Step 2: Configure Extension

1.  Press `Ctrl+Shift+P` (or `Cmd+Shift+P` on Mac).
2.  Run the command: `TakaTime: Setup MongoDB URI`.
3.  Paste your connection string and hit Enter.

**That's it!** TakaTime is now tracking your work. 🎉

## 🔧 Commands

- `TakaTime: Setup MongoDB URI`: Opens the configuration prompt to update your database connection or re-download the binary.

## 📊 How to Display Stats on GitHub

To show the graph on your GitHub profile, you need to set up the **Taka-Report** action in your profile repository.

1.  Go to your GitHub Profile repository (`username/username`).
2.  Create a file: `.github/workflows/takatime.yml`.
3.  Copy the workflow configuration from the [Official Repository](https://github.com/Rtarun3606k/TakaTime).

## 🛡️ Privacy Policy

**TakaTime does not send your data to any third-party servers.**

- Your coding activity is sent **only** to the MongoDB URI you provide.
- The extension downloads a helper binary from the official [TakaTime GitHub Releases](https://github.com/Rtarun3606k/TakaTime/releases).
- No telemetry is collected by the extension author.

## 🔗 Links

- [GitHub Repository](https://github.com/Rtarun3606k/TakaTime)
- [Report an Issue](https://github.com/Rtarun3606k/TakaTime/issues)

---

**Enjoying TakaTime?** ⭐ Star the [repo on GitHub](https://github.com/Rtarun3606k/TakaTime)!
