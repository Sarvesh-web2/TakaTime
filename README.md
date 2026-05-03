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
  <img src="https://img.shields.io/github/downloads/Rtarun3606k/TakaTime/total?style=for-the-badge&color=success" alt="Total Downloads">
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
  <a href="https://github.com/Rtarun3606k/TakaTime">
    <img src="https://img.shields.io/badge/Antigravity-111111?style=for-the-badge&logo=rocket&logoColor=00E5FF">
  </a>

  <br>

  ![Go](https://img.shields.io/badge/Go-65.8%25-00ADD8?style=flat-square&logo=go&logoColor=white)
  ![JavaScript](https://img.shields.io/badge/JavaScript-22.3%25-F7DF1E?style=flat-square&logo=javascript&logoColor=black)
  ![Lua](https://img.shields.io/badge/Lua-11.9%25-000080?style=flat-square&logo=lua&logoColor=white)
  <img src="https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white" alt="MongoDB">
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">

</div>

<br>

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
---

## Official Documentation & Setup

**[Click here to visit the TakaTime Wiki](https://github.com/Rtarun3606k/TakaTime/wiki)** for complete installation guides, database setup (BYODB), dashboard commands, and theme customization.

---

##  Interactive Terminal Dashboard

TakaTime includes a fully interactive, offline-first terminal dashboard directly inside your editor. View your coding stats, language breakdowns, and project times without ever leaving your workflow or opening a browser.

<div align="center">
  <img src="https://github.com/user-attachments/assets/a15288c7-95b2-49f2-8d50-200b087af36c" width="49%" alt="VS Code Dashboard" />
  <img src="https://github.com/user-attachments/assets/5f53508c-4659-46e6-a0ab-ba5d9d537ed4" width="49%" alt="Neovim Dashboard" />
  <p><em>TakaTime Dashboard running locally in Neovim (left) and VS Code (right)</em></p>
</div>

---

##  Features

- **Non-Blocking Architecture:** Engineered in Go with asynchronous concurrency. Data synchronization occurs entirely in the background, ensuring zero latency impact on your editor's performance.
- **Bring Your Own Database (BYODB):** Data is persisted exclusively to your personal MongoDB instance. This ensures complete data ownership with no third-party tracking or subscription fees.
- **Granular Telemetry:** Intelligently tracks and categorizes development activity by project, programming language, and file type without requiring manual configuration.
- **GitHub Profile Integration:** Automatically generate high-resolution statistical charts for your GitHub Profile README via GitHub Actions.

---

##  Editor Compatibility 

TakaTime is cross-platform and editor-agnostic. All plugins share the same core Go binaries for a consistent experience.

| Feature | Neovim | VS Code | Antigravity | OS Support |
| :--- | :--- | :--- | :--- | :--- |
| **Background Sync** | ✓ Supported | ✓ Supported |✓ Supported | Win, Mac, Linux |
| **Terminal Dashboard**| ✓ Supported | ✓ Supported | ✓ Supported | Win, Mac, Linux |
| **Profile Stats** | ✓ Supported | ✓ Supported |✓ Supported | Win, Mac, Linux |
| **Privacy Controls** | ✓ Supported | ⚙ Planned |⚙ Planned | All OS |

---

##  Architecture

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

## Contributors & Community

We welcome pull requests! Whether you want to add support for a new IDE or a new TUI theme, check out our [Contribution Guidelines](https://github.com/Rtarun3606k/TakaTime/blob/main/CONTRIBUTING.md).

<a href="https://github.com/Rtarun3606k/TakaTime/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Rtarun3606k/TakaTime" />
</a>

**License:** MIT License. See `LICENSE` for details.






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