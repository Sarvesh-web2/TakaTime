
<div align="center">

  <img width="120" alt="TakatimeLogo" src="https://github.com/user-attachments/assets/09cf911b-5246-4b13-99e7-7f435a6cde3a" />

  # TakaTime

  **The Open Source, Self-Hosted WakaTime Alternative.**
  <br>
  <i>"Time is what we want most, but what we use worst."</i>

  <br><br>

  <a href="https://github.com/Rtarun3606k/TakaTime/stargazers">
    <img src="https://img.shields.io/github/stars/Rtarun3606k/TakaTime?style=for-the-badge&logo=star&color=ffea00" alt="GitHub Stars">
  </a>
  <a href="https://github.com/Rtarun3606k/TakaTime/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/Rtarun3606k/TakaTime?style=for-the-badge&color=blue" alt="License">
  </a>

  <br>

  <a href="https://github.com/Rtarun3606k/TakaTime">
    <img src="https://img.shields.io/badge/NeoVim-%2357A143.svg?&style=for-the-badge&logo=neovim&logoColor=white" alt="Neovim">
  </a>
  <a href="https://marketplace.visualstudio.com/items?itemName=Rtarun3606k.takatime">
    <img src="https://img.shields.io/badge/VS%20Code-007ACC?style=for-the-badge&logo=visual-studio-code&logoColor=white" alt="VS Code">
  </a>

  <br>

![Go](https://img.shields.io/badge/Go-65.8%25-00ADD8?style=flat-square&logo=go&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-22.3%25-F7DF1E?style=flat-square&logo=javascript&logoColor=black)
![Lua](https://img.shields.io/badge/Lua-11.9%25-000080?style=flat-square&logo=lua&logoColor=white)
  <img src="https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white" alt="MongoDB">
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">

  ![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/Rtarun3606k/TakaTime/example.yml)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/Rtarun3606k/TakaTime)


</div>

<br>

<!--takatime-start-->

<h2 align="center">TakaTime Weekly Report</h2>

<p align="center">
  <img src="./public/taka-time.png" width="100%" alt="Time Stats" /><br/>
  <img src="./public/taka-languages30.png" width="400" alt="Languages" />
  <img src="./public/taka-projects30.png" width="400" alt="Projects" /><br/>
  <img src="./public/taka-languages.png" width="400" alt="Languages" />
  <img src="./public/taka-projects.png" width="400" alt="Projects" /><br/>
  <img src="./public/taka-tech.png" width="100%" alt="Tech Stack" />
</p>

<p align="center"><em>Generated automatically by <a href="https://github.com/Rtarun3606k/TakaTime">TakaTime</a></em></p>

<!--takatime-end-->
---

## Table of Contents
- [Features](#features)
- [How it Works](#How-it-Works)
- [Installation](#installation)
  - [Using VS Code](#using-vs-code)
  - [Using lazy.nvim](#using-lazynvim)
- [Setup Guide](#setup-guide)
- [GitHub Profile Stats Setup](#how-to-add-stats-to-your-github-profile)
  - [Customization & Themes](#Customization-&-Themes)
  - [Configuration Parameters](#Configuration-Parameters)
- [Troubleshooting](#troubleshooting)
- [Disclaimer & Roadmap](#disclaimer--roadmap)
- [License](#license)

---

## Features

- **Non-Blocking Architecture** Engineered in Go with asynchronous concurrency. Data synchronization occurs entirely in the background, ensuring zero latency impact on the editor's performance.
- **Privacy-Centric** Storage Data is persisted exclusively to your personal MongoDB instance. This self-hosted model ensures complete data ownership with no third-party tracking, telemetry, or subscription fees.
- **Automated Dependency Management** The plugin automatically detects the host operating system (Linux/macOS) and retrieves the appropriate pre-compiled binary during the initial setup.
- **Portfolio Visualization** Includes a dedicated CLI utility for generating high-resolution statistical charts, optimized for seamless integration into GitHub Profile READMEs.
- **Granular Telemetry** Intelligently tracks and categorizes development activity by project, programming language, and file type without requiring manual configuration.

---

## How it Works


<div align="center">
  <table border="0">
    <tr>
      <th align="center">High-Level Architecture</th>
      <th align="center">Zero-Latency Flow</th>
    </tr>
    <tr>
      <td width="50%" valign="top">
        <img src="https://github.com/user-attachments/assets/0aa39476-12c7-4cdd-9b27-8e985fcce29d" alt="Sequence Diagram" width="100%">
      </td>
      <td width="50%" valign="top">
        <img src="https://github.com/user-attachments/assets/9a844b2c-d018-470f-9dcc-fabf5a0ac3cf" alt="High Level Architecture Diagram" width="100%">
      </td>
    </tr>
  </table>
</div>

---

## Installation

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

- ### Step 1: Get a Database

  You need a MongoDB connection string. You have two free options:
  - add all ip access `(if you want github Stats its required)`
  - <img width="986" height="752" alt="image" src="https://github.com/user-attachments/assets/d9433977-e841-4e0d-a1d2-9847901501d6" />
  - Cloud (Recommended): Create a free account on MongoDB Atlas. Create a free cluster and get your connection string (e.g., mongodb+srv://user:pass@cluster...).
  - Local (Docker): Run docker run -d -p 27017:27017 mongo.
  
- ### Step 2: Initialize the Plugin
  
  Open Neovim.
  
  Run the setup command:
  Vim Script
  
  ```nvim
  :TakaInit
  ```
  
  Paste your MongoDB Connection String when prompted. (This is saved securely in your local data folder, ~/.local/share/nvim/taka_data.json).
  
- ### Step 3: Verify

  Run the status command to check if everything is working:
  Vim Script
  
  ```nvim
  :TakaStatus
  ```
  
  If it says "TakaTime is configured and running," you are good to go!
  
  ---

## GitHub Profile Stats Setup

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

<h2 align="center">TakaTime Weekly Report</h2>

<p align="center">
  <img src="./public/taka-time.png" width="100%" alt="Time Stats" /><br/>
  <img src="./public/taka-languages30.png" width="400" alt="Languages" />
  <img src="./public/taka-projects30.png" width="400" alt="Projects" /><br/>
  <img src="./public/taka-languages.png" width="400" alt="Languages" />
  <img src="./public/taka-projects.png" width="400" alt="Projects" /><br/>
  <img src="./public/taka-tech.png" width="100%" alt="Tech Stack" />
</p>

<p align="center"><em>Generated automatically by <a href="https://github.com/Rtarun3606k/TakaTime">TakaTime</a></em></p>

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

- ## Customization & Themes

    Taka-Report supports full customization through command-line flags. You can choose from pre-built themes or override specific colors to match your GitHub profile aesthetic.
    
 -  ### 1. Base Themes
    Use the `-theme` flag to apply a pre-configured color palette.  
    **Default:** `dark`
      
| Theme | Description |
| :--- | :--- |
| `dark` | GitHub Dark Dimmed (Default) |
| `light` | GitHub Light |
| `dracula` | Dracula Color Palette |
| `nord` | Nord Winter Color Palette |
| `gruvbox` | Gruvbox Retro |
| `monokai` | Monokai Vivid |
| `cyberpunk` | High Contrast Neon |
| `tokyonight` | Tokyo Night Deep Blue Palette |
| `everforest` | Everforest Soft Nature Theme |
| `iceberg` | Iceberg Cool Minimal Blues |
| `sunset` | Warm Sunset Gradient Colors |
| `deepocean` | Deep Ocean Dark Blue Theme |
| `midnight` | Midnight Purple Developer Theme |
| `catppuccin` | Catppuccin Mocha Pastel Palette |
| `solarized` | Solarized Dark Classic Palette |
| `onedark` | OneDark Pro VSCode Style |
| `material` | Material Dark UI Theme |
| `synthwave` | Retro Synthwave Neon Colors |
    
    **Usage Example:**
    ```bash
    ./taka-report -theme nord
    ```

- ## Configuration Parameters

    You can pass these flags to the `taka-report` binary to control the data scope and visual style of your report.
    
    | Flag | Type | Default | Description |
    | :--- | :--- | :--- | :--- |
    | **`-days`** | `int` | `0` | **Data Scope:** No longer in use just set it to Zero 0|
    | **`-theme`** | `string` | `"dark"` | **Base Theme:** Selects a pre-configured color palette. <br>Options: `dark`, `light`, `dracula`, `nord`, `gruvbox`, `monokai`, `cyberpunk` |
    | **`-bg`** | `hex` | *Theme* | **Background:** Overrides the main card background color. |
    | **`-text`** | `hex` | *Theme* | **Primary Text:** Overrides the color of main headers and key statistics. |
    | **`-subtext`** | `hex` | *Theme* | **Secondary Text:** Overrides the color of labels, timestamps, and axis text. |
    | **`-bar-bg`** | `hex` | *Theme* | **Bar Background:** Overrides the color of the empty/unfilled portion of progress bars. |
    | **`-c1`** | `hex` | *Theme* | **Primary Accent:** Used for the highest values (e.g., "All Time" stat) and primary progress bars. |
    | **`-c2`** | `hex` | *Theme* | **Secondary Accent:** Used for medium-high values (e.g., "Last 30 Days"). |
    | **`-c3`** | `hex` | *Theme* | **Tertiary Accent:** Used for medium-low values (e.g., "Last 7 Days"). |
    | **`-c4`** | `hex` | *Theme* | **Quaternary Accent:** Used for the lowest values (e.g., "Yesterday") or distinct highlights. |
    
    > **Note:** Color overrides (like `-bg`) take precedence over the base `-theme`. You can start with `-theme dracula` and then just change the background with `-bg "#000000"`.

    Exapmle use  Neon Theme :
  ```bash
  ./taka-report -days=7 -bg "#0d1117" -text "#00FF00" -subtext "#008800" -bar-bg "#111111" -c1 "#00FF00" -c2 "#00DD00" -c3 "#00AA00" -c4 "#005500"
  ```

---

## Troubleshooting

"TakaTime is not configured"

    Run :TakaInit again and ensure your URI is correct.

    Check if the secret file exists: ~/.local/share/nvim/taka_data.json.

Upload Failed / Syncing Forever

Enable debug mode in your config:

```Lua

 return {
  "Rtarun3606k/TakaTime",
  lazy = false,
  config = function()
    -- Optional: Enable debug mode if you run into issues
    require("taka-time").setup({
        debug = true
    })
  end,
}
```

Run :messages in Neovim to see the logs.

Ensure your IP address is whitelisted in MongoDB Atlas.

## Advanced Debugging (Log File)

If you are still facing issues, TakaTime maintains a persistent log file that tracks all binary operations, network requests, and errors.

Log Location:

    Linux/macOS:  ~/.takatime/debug-logs.log

    Windows:  C:\Users\<YourUser>\.takatime\debug-logs.log

Reporting an Issue: If you open a GitHub Issue, please attach this log file (or paste the last 50 lines). It helps us fix bugs 10x faster!


---

For Changes Look for `CHANGELOG.md`

---

## Disclaimer & Roadmap

`Active Beta`: This project is currently in active development. We might introduce breaking changes in future updates as we refine the architecture. Use at your own risk.

`Documentation`: This documentation was generated with the assistance of AI. While we strive for accuracy, there may be minor errors or typos.

`Feedback`: If you encounter any bugs, have feature requests, or notice documentation errors, please feel free to open an issue or report it to **Rtarun3606k**.


---

## License

MIT License. See `LICENSE` for details.

