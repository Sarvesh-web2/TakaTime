## Relese : [2.0.5]NeoVim/upload/reporter [0.0.3] VsCode - 2026-01-26 (The debugging and New Metrics Update)

### Added

- **Cross-Platform Logging:** Implemented a robust file logger that automatically stores debug logs in `~/.takatime/debug-logs.log` (works on Windows, Linux, and macOS).
- **Editor Tracking:** Added a new `-editor` flag to the CLI and an `editor` field to the log schema. You can now distinguish between coding sessions in Neovim vs. VS Code.
- **Marketplace SEO:** Added keywords, categories, and a repository link to the VS Code extension `package.json` to improve search visibility.

### Changed

- **Git Context Logic:** Updated `GetGitBranch` to run git commands in the _file's_ directory rather than the _process_ directory. This fixes issues where the branch was not detected in multi-root workspaces.
- **Extension Metadata:** Expanded the extension description to better highlight features like metrics and privacy.

### Fixed

- **Directory Safety:** The logger now automatically creates the `.takatime` directory if it does not exist, preventing startup crashes on fresh installs.

## 🚀 Release: v2.0.3-beta (The "User Experience" Update)

**Description:** This is a massive update focused on User Experience and Ease of Installation. We have completely removed the need to hardcode secrets in your Lua config.
The plugin now handles binary management automatically,
downloading the correct tools for your OS upon installation.

**Changelog:**

- feat(core): add :TakaInit command for secure, interactive setup.
- feat(install): implement auto-download logic for taka-upload and taka-report binaries.
- feat(storage): move secret storage to stdpath("data") (secure JSON file) instead of init.lua.
- fix(ci): update release workflow to build and upload taka-report binary (fixes "asset not found" error).
- fix(ui): silence "Syncing..." messages by default (set debug=false in config).
- refactor: split logic into core, storage, and utils modules for better maintainability.

---

## 📦 Release: v1.0.1 (The "Foundation" Release)

**Description:** The first stable release of TakaTime.nvim. This version lays the groundwork for privacy-focused, self-hosted time tracking. It connects Neovim directly to your MongoDB instance using a
high-performance Go binary.

**Chnage logs**

- feat: initial release of Lua plugin structure.
- feat: implement Go binary for MongoDB uploads.
- feat: add debounce logic to prevent spamming the database on every keystroke.
- config: basic setup function with mongo_uri configuration support.
