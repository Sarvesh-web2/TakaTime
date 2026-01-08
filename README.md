# ⏳ TakaTime.nvim

> "Time is what we want most, but what we use worst."

**TakaTime** is a blazingly fast, privacy-focused coding time tracker for Neovim.

It works just like WakaTime, but with one major difference: **You own your data.**
Instead of sending your coding activity to a third-party server, TakaTime stores everything in **your own MongoDB database**.

## ✨ Features

- 🚀 **Zero Latency:** Written in Go. Uploads happen asynchronously in the background so it never blocks your typing.
- 🔒 **Privacy First:** Data is stored in your personal MongoDB (Free Tier on Atlas). No subscriptions, no tracking.
- 📦 **Auto-Install:** automatically downloads the correct binary for your OS (Linux/Mac) on first run.
- 📊 **GitHub Profile Stats:** Includes a CLI tool to generate beautiful charts for your GitHub Profile README.
- 📂 **Smart Tracking:** Tracks Projects, Languages, and Files automatically.

---

## 📦 Installation

### Using [lazy.nvim](https://github.com/folke/lazy.nvim)

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

## ⚙️ Setup Guide

Step 1: Get a Database

You need a MongoDB connection string. You have two free options:

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
> _Jan 01_ to _Jan 08_

> [!TIP]
> **Total Coding Time (7d):** 4h 18m

#### 📈 Trends
| Period        | Duration    | Period       | Duration    |
| :---          | :---        | :---         | :---        |
| Yesterday     | **35m**     | Last 7 Days  | **4h 18m**  |
| Last 30 Days  | **4h 18m**  | All Time     | **4h 18m**  |

#### 💻 Languages
| Language | Time | Percentage |
| :--- | :--- | :--- |
| **go** | 1h 42m | 🟦🟦🟦⬜⬜⬜⬜⬜⬜⬜ 39.7% |
| **lua** | 1h 3m | 🟦🟦⬜⬜⬜⬜⬜⬜⬜⬜ 24.7% |
| **txt** | 41m | 🟦⬜⬜⬜⬜⬜⬜⬜⬜⬜ 15.9% |
| **text** | 27m | 🟦⬜⬜⬜⬜⬜⬜⬜⬜⬜ 10.6% |
| **Other** | 23m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 9.1% |

#### 🔥 Projects
| Project | Time | Percentage |
| :--- | :--- | :--- |
| **taka-time.nvim** | 2h 37m | 🟩🟩🟩🟩🟩🟩⬜⬜⬜⬜ 60.7% |
| **testTakaTime** | 41m | 🟩⬜⬜⬜⬜⬜⬜⬜⬜⬜ 15.9% |
| **nvim** | 24m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 9.6% |
| **vscodePlugin** | 23m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 9.1% |
| **Other** | 12m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 4.7% |


<!--takatime-end-->
```

3. Create the Workflow

Create a file in your repo at .github/workflows/update-stats.yml and paste this content:

```yml
name: Update TakaTime Stats

on:
  schedule:
    - cron: "0 0 * * *" # Runs every midnight UTC
  workflow_dispatch: # Allows manual trigger

jobs:
  update-readme:
    runs-on: ubuntu-latest
    permissions:
      contents: write

    steps:
      - name: Download Taka-Report Binary
        env:
          GH_TOKEN: ${{ github.token }}
        run: |
          gh release download --repo Rtarun3606k/TakaTime --pattern "taka-report-linux-amd64" --output taka-report
          chmod +x taka-report

      - name: Generate Report & Update Profile
        env:
          MONGO_URI: ${{ secrets.MONGO_URI }}
          TARGET_REPO: ${{ github.repository }} # Automatically targets this repo
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
> _Jan 01_ to _Jan 08_

> [!TIP]
> **Total Coding Time (7d):** 4h 18m

#### 📈 Trends
| Period        | Duration    | Period       | Duration    |
| :---          | :---        | :---         | :---        |
| Yesterday     | **35m**     | Last 7 Days  | **4h 18m**  |
| Last 30 Days  | **4h 18m**  | All Time     | **4h 18m**  |

#### 💻 Languages
| Language | Time | Percentage |
| :--- | :--- | :--- |
| **go** | 1h 42m | 🟦🟦🟦⬜⬜⬜⬜⬜⬜⬜ 39.7% |
| **lua** | 1h 3m | 🟦🟦⬜⬜⬜⬜⬜⬜⬜⬜ 24.7% |
| **txt** | 41m | 🟦⬜⬜⬜⬜⬜⬜⬜⬜⬜ 15.9% |
| **text** | 27m | 🟦⬜⬜⬜⬜⬜⬜⬜⬜⬜ 10.6% |
| **Other** | 23m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 9.1% |

#### 🔥 Projects
| Project | Time | Percentage |
| :--- | :--- | :--- |
| **taka-time.nvim** | 2h 37m | 🟩🟩🟩🟩🟩🟩⬜⬜⬜⬜ 60.7% |
| **testTakaTime** | 41m | 🟩⬜⬜⬜⬜⬜⬜⬜⬜⬜ 15.9% |
| **nvim** | 24m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 9.6% |
| **vscodePlugin** | 23m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 9.1% |
| **Other** | 12m | ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜ 4.7% |


<!--takatime-end-->
