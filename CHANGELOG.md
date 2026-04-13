# Changelog

All notable changes to the **TakaTime** project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

### [2.2.2] -Release v2.2.2 (The Antigravity Patch)

This minor patch cleans up how the dashboard processes background data, officially making room to display new editors!

### What's New
* **Antigravity Tracking Supported:** The dashboard now properly displays time tracked in the `antigravity` editor!
* **Data Cleanup:** Filtered out `unknown` editor data from the local UI, ensuring your actual top editors accurately claim their spots on the leaderboard.

---
###  How to Install Manually (Using `.vsix`)
1. **Download the file:** [Click here to download takatime-0.1.1.vsix](https://github.com/Rtarun3606k/TakaTime/releases/download/v2.2.2/takatime-0.1.1.vsix) directly.
3. **Install from VSIX:** Click the **...** (three dots/gear icon) at the top right of the Extensions panel and select **"Install from VSIX..."**.
4. **Select the file:** Locate and select the `.vsix` file you just downloaded.
5. **Configure Database:** Once installed, run the `TakaTime: Setup` command (or click the status bar) and enter your MongoDB Connection String.

---

## [2.2.0] - 2026-04-09

### New Features
* **Interactive Dashboard:** Introduced a fully responsive, terminal-based dashboard powered by Go and Bubble Tea to view local stats without a browser.
* **VS Code Dashboard Integration:** Added `TakaTime: Open Dashboard` command and a dedicated quick-access Graph Icon in the editor title menu. The dashboard opens cleanly within the main editor area.
* **Neovim Dashboard Integration:** Added the `:TakaDash` command to launch the interactive UI in a centered, floating terminal window.
* **Multi-Binary Auto-Downloader:** The VS Code extension now intelligently downloads and manages multiple required binaries (`taka-upload` and `taka-dashboard`) with a unified progress UI.

###  Architecture & Refactoring
* **Modularized Go Binaries:** Split the monolithic architecture into purpose-built, lightweight binaries (`taka-upload`, `taka-dashboard`, and `taka-report`) to improve performance, reduce memory footprint, and allow for independent updates.
* **Decoupled Configuration:** Refactored the internal Node.js extension architecture to isolate the `Uploader`, `Downloader`, `HeartBeat`, and `Config` modules for better maintainability.
* **Explicit Payload Routing:** Updated the heartbeat spawn process to pass the MongoDB URI directly via command-line flags, improving reliability during background sync.

###  Improvements & Fixes
* **VS Code UI Optimization:** Adjusted terminal creation logic so the Bubble Tea layout renders perfectly without stretching on ultra-wide monitors.
* **Status Bar Clarity:** Updated the status bar to reflect the health of *all* required binaries, actively warning users if core files are missing.
* **Safe Process Spawning:** Enhanced the background Go process to ensure it detaches cleanly from VS Code, preventing zombie processes or editor lag.

---

## [2.1.0] - 2026-02-12

### Added
- **High-DPI Support:** Increased canvas resolution to 1200px width for all cards to ensure crisp rendering on GitHub Retina displays.
- **Theming Engine:** Added `types.ThemeConfig` to support custom color palettes.
- **Light Mode:** Implemented a GitHub-style Light Theme palette (`types.LightMode`).
- **Truncation Helper:** Added logic to shorten long project/language names (e.g., "taka-tim..") to prevent layout breaking.

### Changed
- **Time Card Layout:** Completely redesigned from a "2x2 Grid" to a sleek "Horizontal Strip" (120px height) to save vertical space.
- **Tech Stack Layout:**
  - Increased gap between "Editors" and "Operating Systems" columns to prevents text collision.
  - Shortened progress bars (300px) to allow more breathing room for labels.
- **Fonts:** Bumped font sizes across the board (Header: 65px, Data: 42px) for better readability.

### Fixed
- **Overlap Issues:** Fixed a bug where the "Coding Activity" header would overlap with data columns by moving the title to the top-left corner.
- **Data Noise:** Added filters to hide "unknown" editors and operating systems from the environment stats.
- **Zero-Value Visibility:** Added minimum width logic (`20px`) to progress bars so tools with 0% usage (but tracked) remain visible as a thin strip.

## [0.1.0] - 2025-12-31
### Added
- Initial release of TakaTime reporter.
- Basic image generation for Languages and Projects.
- MongoDB connection logic.

---

## [2.0.5] - 2026-01-26
### Added
- **Cross-Platform Logging:** Implemented a robust file logger that stores debug logs in `~/.takatime/debug-logs.log` (Windows, Linux, macOS).
- **Editor Tracking:** Added `-editor` flag and `editor` field to log schema to distinguish between Neovim and VS Code sessions.
- **Marketplace SEO:** Added keywords, categories, and repository links to the VS Code extension to improve visibility.

### Changed
- **Git Context:** Updated `GetGitBranch` to run git commands in the *file's* directory instead of the process directory, fixing branch detection in multi-root workspaces.
- **Metadata:** Expanded extension description to highlight new metrics and privacy features.

### Fixed
- **Directory Safety:** The logger now automatically creates the `.takatime` directory if missing, preventing startup crashes on fresh installs.

---

## [2.0.3-beta]
### Added
- **Interactive Setup:** Added `:TakaInit` command for secure, interactive configuration.
- **Auto-Installation:** Implemented logic to automatically download `taka-upload` and `taka-report` binaries for the user's OS.
- **Secure Storage:** Moved secret storage to `stdpath("data")` (secure JSON) instead of requiring hardcoded secrets in `init.lua`.

### Changed
- **Refactoring:** Split logic into core, storage, and utils modules for better maintainability.

### Fixed
- **CI Workflow:** Updated release workflow to correctly build and upload the `taka-report` binary (fixes "asset not found" errors).
- **UI UX:** Silenced "Syncing..." messages by default (controlled via `debug=false` config).

---

## [1.0.1] - The "Foundation" Release
### Added
- **Initial Release:** Launched the first stable version of the Lua plugin structure.
- **Go Integration:** Implemented high-performance Go binary for direct MongoDB uploads.
- **Debouncing:** Added logic to prevent database spamming on every keystroke.
- **Configuration:** Added basic setup function with `mongo_uri` support.
