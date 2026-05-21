#  TakaTime

![Version](https://img.shields.io/badge/version-v2.2.4-blue.svg)
![VS Code](https://img.shields.io/badge/VS%20Code-Supported-blueviolet)
![JetBrains](https://img.shields.io/badge/JetBrains-Supported-orange)
![License](https://img.shields.io/badge/license-MIT-green)

**An open-source, privacy-first coding telemetry engine and time tracker.**

TakaTime seamlessly tracks your coding activity across multiple IDEs and provides a beautiful Terminal User Interface (TUI) dashboard right inside your editor. No third-party servers, no subscription fees—your data stays completely yours in your own MongoDB instance.

---

## Features

* **Cross-Platform Parity:** Works flawlessly whether you are coding in VS Code, IntelliJ IDEA, PyCharm, WebStorm, or GoLand.
* **Intelligent Heartbeat:** Accurately calculates true coding duration and automatically pauses when you step away from the keyboard.
* **Integrated TUI Dashboard:** View your coding analytics natively inside your IDE's built-in terminal—no context switching required.
* **Auto-Updating Ecosystem:** Smart semantic version checking ensures your background Go binaries stay perfectly up to date without breaking your workflow.
* **100% Privacy:** Bring your own database (BYOD). All telemetry is routed directly to your personal MongoDB.

---

##  Installation

## Official Documentation & Setup

**[Click here to visit the TakaTime Wiki](https://github.com/Rtarun3606k/TakaTime/wiki)** for complete installation guides, database setup (BYODB), dashboard commands, and theme customization.



### 1. Visual Studio Code
1. Open VS Code and navigate to the Extensions tab (`Ctrl+Shift+X`).
2. Search for **TakaTime** and click Install.
3. *Alternatively: Install directly from the [VS Code Marketplace](#).*

### 2. JetBrains IDEs (IntelliJ, PyCharm, etc.)
1. Open your IDE Settings (`Ctrl+Alt+S` or `Cmd+,`).
2. Navigate to **Plugins** > **Marketplace**.
3. Search for **TakaTime** and click Install.
4. *Alternatively: Install directly from the [JetBrains Marketplace](#).*

---

## Quick Start Setup

1. **Get a MongoDB URI:** Spin up a free cluster on MongoDB Atlas or run one locally via Docker.
2. **Configure the Plugin:** * Look at the bottom Status Bar of your IDE and click **TakaTime**.
  * Select **⚙️ Configure MongoDB URI** from the popup menu.
  * Paste your URI and hit Enter.
3. **Automatic Download:** TakaTime will instantly download the highly optimized Go binaries (`taka-upload` and `taka-dashboard`) in the background.
4. **Start Coding:** Just start typing. TakaTime will silently log your language, project name, and duration!

---

## The Dashboard

Want to see your stats? Just click **TakaTime** in your Status Bar and select **🚀 Open Dashboard**.

TakaTime will automatically spawn a new terminal instance inside your IDE and boot up the Go TUI, giving you immediate insights into your coding habits, favorite languages, and active projects.

---

##  Architecture Overview

TakaTime is built on a highly modular, decoupled architecture to ensure zero IDE UI freezing:
* **The Clients:** Lightweight TypeScript (VS Code) and Kotlin (JetBrains) plugins that act purely as event listeners and UI layers.
* **The Engine:** A fast, compiled Go binary (`taka-upload`) that handles all asynchronous network requests to MongoDB.
* **The Visualizer:** A Go-based Terminal User Interface (`taka-dashboard`) for local analytics.

##  Editor Compatibility

TakaTime is cross-platform and editor-agnostic. All plugins share the same core Go binaries for a consistent experience.

<div align="center">

<table>
  <tr>
    <th>Feature</th>
    <th>Neovim</th>
    <th>VS Code</th>
    <th>Antigravity</th>
    <th>JetBrains</th>
    <th>OS Support</th>
  </tr>
  <tr>
    <td><b>Background Sync</b></td>
    <td>✓ Supported</td>
    <td>✓ Supported</td>
    <td>✓ Supported</td>
    <td>✓ Supported</td>
    <td>Win, Mac, Linux</td>
  </tr>
  <tr>
    <td><b>Terminal Dashboard</b></td>
    <td>✓ Supported</td>
    <td>✓ Supported</td>
    <td>✓ Supported</td>
    <td>✓ Supported</td>
    <td>Win, Mac, Linux</td>
  </tr>
  <tr>
    <td><b>Profile Stats</b></td>
    <td>✓ Supported</td>
    <td>✓ Supported</td>
    <td>✓ Supported</td>
    <td>✓ Supported</td>
    <td>Win, Mac, Linux</td>
  </tr>
  <tr>
    <td><b>Privacy Controls</b></td>
    <td>✓ Supported</td>
    <td>⚙ Planned</td>
    <td>⚙ Planned</td>
    <td>⚙ Planned</td>
    <td>All OS</td>
  </tr>
</table>

</div>



## Contributing
Contributions are always welcome! Whether it's adding support for Neovim, creating new dashboard widgets, or optimizing the Go binaries.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

##  License
Distributed under the MIT License. See `LICENSE` for more information.