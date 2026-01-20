<div align="center">
  <table border="0" cellspacing="0" cellpadding="0">
    <tr border="0">
      <td valign="middle" align="center" border="0">
        <img width="200" alt="TakatimeLogo" src="https://github.com/user-attachments/assets/09cf911b-5246-4b13-99e7-7f435a6cde3a" />
      </td>
      <td valign="middle" border="0">
        <h1>TakaTime.nvim</h1>
      </td>
    </tr>
  </table>
</div>

<p align="center">
  <img src="https://github.com/Rtarun3606k/TakaTime/blob/main/public/dashboard-preview-vscode.png" alt="TakaTime Banner" width="100%">
</p>

> "Time is what we want most, but what we use worst."

**TakaTime** is a blazingly fast, privacy-focused coding time tracker for Neovim.

It works just like WakaTime, but with one major difference: **You own your data.**
Instead of sending your coding activity to a third-party server, TakaTime stores everything in **your own MongoDB database**.

---

## 📖 Table of Contents
- [✨ Features](#-features)
- [📦 Installation](#-installation)
  - [Using VS Code](#using-vs-code)
  - [Using lazy.nvim](#using-lazynvim)
- [⚙️ Setup Guide](#-setup-guide)
- [📊 GitHub Profile Stats](#-how-to-add-stats-to-your-github-profile)
- [🛠️ Troubleshooting](#-troubleshooting)
- [⚠️ Disclaimer & Roadmap](#-disclaimer--roadmap)
- [📄 License](#-license)

---

## ✨ Features

- 🚀 **Zero Latency:** Written in Go. Uploads happen asynchronously in the background so it never blocks your typing.
- 🔒 **Privacy First:** Data is stored in your personal MongoDB (Free Tier on Atlas). No subscriptions, no tracking.
- 📦 **Auto-Install:** automatically downloads the correct binary for your OS (Linux/Mac) on first run.
- 📊 **GitHub Profile Stats:** Includes a CLI tool to generate beautiful charts for your GitHub Profile README.
- 📂 **Smart Tracking:** Tracks Projects, Languages, and Files automatically.

## 🏗️ How it Works


<div align="center">
  <table border="0">
    <tr>
      <th align="center">High-Level Architecture</th>
      <th align="center">Zero-Latency Flow</th>
    </tr>
    <tr>
      <td width="50%" valign="top">
        <img src="https://github.com/user-attachments/assets/edee6d78-034e-4f95-a0e0-4a1616180f1d" alt="Sequence Diagram" width="100%">
      </td>
      <td width="50%" valign="top">
        <img src="https://github.com/user-attachments/assets/37420b31-e5e4-4ff0-a823-84261db1c5a6" alt="High Level Architecture Diagram" width="100%">
      </td>
    </tr>
  </table>
</div>

---

## 📦 Installation

### Using VS Code 

https://github.com/user-attachments/assets/a3c492d8-898c-497a-bc0c-c2f8ebc5d03b

---

### Using [lazy.nvim](https://github.com/folke/lazy.nvim)



https://github.com/user-attachments/assets/edf09531-ed66-4709-9b78-5edc90843510



Add this to your plugin configuration:

```lua
return {
  "Rtarun3606k/TakaTime",
  lazy = false,
  config = function()
    -- Optional: Enable debug mode if you run into issues
    require("taka-time").setup({
        debug = false
    })
  end,
}
```

---

##  Setup Guide

Step 1: Get a Database

You need a MongoDB connection string. You have two free options:
- add all ip access `(only if you keeping changing wifi)`
- <img width="986" height="752" alt="image" src="https://github.com/user-attachments/assets/d9433977-e841-4e0d-a1d2-9847901501d6" />
- Cloud (Recommended): Create a free account on MongoDB Atlas. Create a free cluster and get your connection string (e.g., mongodb+srv://user:pass@cluster...).
- Local (Docker): Run docker run -d -p 27017:27017 mongo.

Step 2: Initialize the Plugin

Open Neovim.

Run the setup command:
Vim Script

```nvim
:TakaInit
```

Paste your MongoDB Connection String when prompted. (This is saved securely in your local data folder, ~/.local/share/nvim/taka_data.json).

Step 3: Verify

Run the status command to check if everything is working:
Vim Script

```nvim
:TakaStatus
```

If it says "✅ TakaTime is configured and running," you are good to go!

---

## 📊 How to Add Stats to Your GitHub Profile

TakaTime comes with a report generator that works with GitHub Actions to update your Profile README automatically.

1. Prepare your Profile Repo

   Go to your GitHub Profile Repository (the one named username/username).

   Go to Settings > Secrets and variables > Actions.

   Add a New Repository Secret named MONGO_URI with your connection string.

   (Optional) Add `GIST_TOKEN` if you plan to use Gists (not required for direct README updates).

2. Add the Markers

- Add start and end markers to your README.md

```md
<!--takatime-start-->

> [!NOTE]
> **TakaTime Dashboard**
> _Jan 13_ to _Jan 20_

> [!TIP]
> **Total Coding Time (7d):** 2h 29m

#### 📈 Trends
| Period        | Duration    | Period       | Duration    |
| :---          | :---        | :---         | :---        |
| Yesterday     | **0m**      | Last 7 Days  | **2h 29m**  |
| Last 30 Days  | **6h 48m**  | All Time     | **6h 48m**  |

#### 💻 Languages
| Language | Time | Percentage |
| :--- | :--- | :--- |
| **go** | 1h 9m | 🟦🟦🟦🟦⬜⬜⬜⬜⬜⬜ 46.3% |
| **lua** | 1h 6m | 🟦🟦🟦🟦⬜⬜⬜⬜⬜⬜ 44.3% |
| **json** | 6m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 4.0% |
| **markdown** | 4m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 2.7% |
| **Other** | 4m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 2.7% |

#### 🔥 Projects
| Project | Time | Percentage |
| :--- | :--- | :--- |
| **testTakaTime** | 1h 9m | 🟩🟩🟩🟩⬜⬜⬜⬜⬜⬜ 46.3% |
| **plugins** | 1h 5m | 🟩🟩🟩🟩⬜⬜⬜⬜⬜⬜ 43.6% |
| **Takatime** | 14m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 9.4% |
| **taka-time.nvim** | 1m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 0.7% |


<!--takatime-end-->
```

3. Create the Workflow

Create a file in your repo at .github/workflows/update-stats.yml and paste this content:

```yml

name: Update TakaTime Stats

on:
  schedule:
    - cron: "0 0 * * *" # Runs every midnight UTC
  workflow_dispatch:      # Allows manual trigger

jobs:
  update-readme:
    runs-on: ubuntu-latest
    permissions:
      contents: write # Needed to download releases

    steps:
      - name: Download Taka-Report Binary
        env:
          GH_TOKEN: ${{ github.token }}
        run: |
          # Downloads the latest stable binary
          gh release download --repo Rtarun3606k/TakaTime --pattern "taka-report-linux-amd64" --output taka-report
          chmod +x taka-report

      - name: Generate Report & Update Profile
        env:
          MONGO_URI: ${{ secrets.MONGO_URI }}
          GIST_TOKEN: ${{ github.token }}
          TARGET_REPO: ${{ github.repository }}
        run: ./taka-report -days=7
```

`Note:` This workflow downloads the taka-report tool and runs it against your database to generate stats.

---

## 🛠️ Troubleshooting

"TakaTime is not configured"

    Run :TakaInit again and ensure your URI is correct.

    Check if the secret file exists: ~/.local/share/nvim/taka_data.json.

Upload Failed / Syncing Forever

Enable debug mode in your config:

```Lua

   require("taka-time").setup({ debug = true })
```

Run :messages in Neovim to see the logs.

Ensure your IP address is whitelisted in MongoDB Atlas.

For Changes Look for `CHANGELOG.md`

---

## ⚠️ Disclaimer & Roadmap

`Active Beta`: This project is currently in active development. We might introduce breaking changes in future updates as we refine the architecture. Use at your own risk.

`Documentation`: This documentation was generated with the assistance of AI. While we strive for accuracy, there may be minor errors or typos.

`Feedback`: If you encounter any bugs, have feature requests, or notice documentation errors, please feel free to open an issue or report it to **Rtarun3606k**.

New screenshots and visual updates will be added soon!

---

## 📄 License

MIT License. See `LICENSE` for details.

<!--takatime-start-->

> [!NOTE]
> **TakaTime Dashboard**
> _Jan 13_ to _Jan 20_

> [!TIP]
> **Total Coding Time (7d):** 2h 29m

#### 📈 Trends
| Period        | Duration    | Period       | Duration    |
| :---          | :---        | :---         | :---        |
| Yesterday     | **0m**      | Last 7 Days  | **2h 29m**  |
| Last 30 Days  | **6h 48m**  | All Time     | **6h 48m**  |

#### 💻 Languages
| Language | Time | Percentage |
| :--- | :--- | :--- |
| **go** | 1h 9m | 🟦🟦🟦🟦⬜⬜⬜⬜⬜⬜ 46.3% |
| **lua** | 1h 6m | 🟦🟦🟦🟦⬜⬜⬜⬜⬜⬜ 44.3% |
| **json** | 6m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 4.0% |
| **markdown** | 4m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 2.7% |
| **Other** | 4m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 2.7% |

#### 🔥 Projects
| Project | Time | Percentage |
| :--- | :--- | :--- |
| **testTakaTime** | 1h 9m | 🟩🟩🟩🟩⬜⬜⬜⬜⬜⬜ 46.3% |
| **plugins** | 1h 5m | 🟩🟩🟩🟩⬜⬜⬜⬜⬜⬜ 43.6% |
| **Takatime** | 14m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 9.4% |
| **taka-time.nvim** | 1m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 0.7% |


<!--takatime-end-->
