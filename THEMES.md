# 🎨 TakaTime Theme Gallery

A visual reference for all built-in TakaTime color themes. Use the theme name/ID in your command or workflow to apply it.

> **How to use a theme:**
> ```bash
> # Using a named theme
> ./taka-report -theme=dracula
>
> # Using custom hex colors (for themes not in the named list)
> ./taka-report -bg "#1f1d2b" -text "#f8f8f2" -subtext "#a599e9" -bar-bg "#2a273f" -c1 "#ff9e64" -c2 "#ffd580" -c3 "#ff6b6b" -c4 "#c678dd"
> ```

---

## Available Themes

| Theme | ID / String | Description |
|-------|-------------|-------------|
| [Dark](#-dark-default) | `dark` | Matrix-inspired green on near-black |
| [Light](#-light) | `light` | Solarized Light — soft cream with blue accents |
| [Dracula](#-dracula) | `dracula` | Classic Dracula — purple highlights on dark grey |
| [Monokai](#-monokai) | `monokai` | Vibrant green & pink on olive-dark background |
| [Cyberpunk](#-cyberpunk) | `cyberpunk` | Neon green & hot pink on pure black |
| [Tokyo Night](#-tokyo-night) | `tokyonight` | Cool blues on deep navy |
| [Catppuccin](#-catppuccin) | `catppuccin` | Pastel blues & pinks on soft dark (Mocha variant) |
| [Synthwave](#-synthwave) | `synthwave` | Retro 80s cyan & hot pink on deep purple |
| [Sunset](#-sunset) | `sunset` | Warm amber & orange on dark background |
| [Midnight Purple](#-midnight-purple) | `midnightpurple` | Deep purple & lavender on very dark background |

---

## 🌑 Dark (Default)

**ID:** `dark` &nbsp;|&nbsp; **Palette:** Bright green (`#39d353`) on near-black (`#0d1117`) — inspired by GitHub's contribution graph

| Stat Card Preview | Terminal Dashboard |
|:-:|:-:|
| ![Dark theme stat card](./public/dark-web.png) | ![Dark theme TUI](./public/dark-tui.png) |

---

## ☀️ Light

**ID:** `light` &nbsp;|&nbsp; **Palette:** Soft cream background (`#FDF6E3`) with blue (`#268BD2`), olive, and magenta accents — Solarized Light inspired

| Stat Card Preview | Terminal Dashboard |
|:-:|:-:|
| ![Light theme stat card](./public/light-web.png) | ![Light theme TUI](./public/light-tui.png) |

---

## 🧛 Dracula

**ID:** `dracula` &nbsp;|&nbsp; **Palette:** Soft purple (`#bd93f9`) and cyan (`#8be9fd`) on dark grey-purple (`#282a36`) — the classic Dracula color scheme

| Stat Card Preview | Terminal Dashboard |
|:-:|:-:|
| ![Dracula theme stat card](./public/dracula-web.png) | ![Dracula theme TUI](./public/dracula-tui.png) |

---

## 🎨 Monokai

**ID:** `monokai` &nbsp;|&nbsp; **Palette:** Vibrant green (`#a6e22e`), hot pink (`#f92672`), and cyan (`#66d9ef`) on an olive-dark background (`#272822`)

| Stat Card Preview | Terminal Dashboard |
|:-:|:-:|
| ![Monokai theme stat card](./public/monokai-web.png) | ![Monokai theme TUI](./public/monokai-tui.png) |

---

## ⚡ Cyberpunk

**ID:** `cyberpunk` &nbsp;|&nbsp; **Palette:** Neon green (`#00ff9f`) and hot pink (`#f6019d`) on near-black (`#0b0e14`) — maximum contrast, maximum energy

| Stat Card Preview | Terminal Dashboard |
|:-:|:-:|
| ![Cyberpunk theme stat card](./public/cyberpunk-web.png) | ![Cyberpunk theme TUI](./public/cyberpunk-tui.png) |

---

## 🌃 Tokyo Night

**ID:** `tokyonight` &nbsp;|&nbsp; **Palette:** Cool blue (`#7aa2f7`) and green (`#9ece6a`) on deep navy (`#1a1b26`) — calm and focused, like coding at midnight in Tokyo

| Stat Card Preview | Terminal Dashboard |
|:-:|:-:|
| ![Tokyo Night theme stat card](./public/tokyonight-web.png) | ![Tokyo Night theme TUI](./public/tokyonight-tui.png) |

---

## 🐱 Catppuccin

**ID:** `catppuccin` &nbsp;|&nbsp; **Palette:** Pastel blue (`#89b4fa`), green (`#a6e3a1`), and pink (`#f38ba8`) on a soft dark background (`#1e1e2e`) — Mocha variant

| Stat Card Preview | Terminal Dashboard |
|:-:|:-:|
| ![Catppuccin theme stat card](./public/catppuccin-web.png) | ![Catppuccin theme TUI](./public/catppuccin-tui.png) |

---

## 🌊 Synthwave

**ID:** `synthwave` &nbsp;|&nbsp; **Palette:** Retro cyan (`#36f9f6`) and hot pink (`#ff5c8a`) on deep purple (`#241b2f`) — pure 1980s nostalgia

| Stat Card Preview | Terminal Dashboard |
|:-:|:-:|
| ![Synthwave theme stat card](./public/synthwave-web.png) | ![Synthwave theme TUI](./public/synthwave-tui.png) |

---

## 🌅 Sunset

**ID:** `sunset` &nbsp;|&nbsp; **Palette:** Warm amber (`#ff9e64`) and gold (`#ffd580`) on a dark plum background (`#1f1d2b`) — golden hour vibes

> **Note:** Use custom hex flags for this theme:
> ```bash
> ./taka-report -bg "#1f1d2b" -text "#f8f8f2" -subtext "#a599e9" -bar-bg "#2a273f" -c1 "#ff9e64" -c2 "#ffd580" -c3 "#ff6b6b" -c4 "#c678dd"
> ```

| Stat Card Preview | Terminal Dashboard |
|:-:|:-:|
| ![Sunset theme stat card](./public/sunset-web.png) | ![Sunset theme TUI](./public/sunset-tui.png) |

---

## 🌙 Midnight Purple

**ID:** `midnightpurple` &nbsp;|&nbsp; **Palette:** Soft purple (`#c084fc`) and blue (`#60a5fa`) on very dark purple (`#1b1325`) — elegant and mysterious

> **Note:** Use custom hex flags for this theme:
> ```bash
> ./taka-report -bg "#1b1325" -text "#e9d8fd" -subtext "#9f7aea" -bar-bg "#2d1b3f" -c1 "#c084fc" -c2 "#60a5fa" -c3 "#34d399" -c4 "#f472b6"
> ```

| Stat Card Preview | Terminal Dashboard |
|:-:|:-:|
| ![Midnight Purple theme stat card](./public/midnightpurple-web.png) | ![Midnight Purple theme TUI](./public/midnightpurple-tui.png) |

---

## Additional Themes

TakaTime supports more themes via custom hex color flags. The following themes are available in the web generator and TUI but use custom colors rather than a named ID:

| Theme | Colors |
|-------|--------|
| Nord | `-bg "#2e3440" -text "#d8dee9" -subtext "#4c566a" -bar-bg "#3b4252" -c1 "#88c0d0" -c2 "#a3be8c" -c3 "#81a1c1" -c4 "#bf616a"` |
| Gruvbox | `-bg "#282828" -text "#ebdbb2" -subtext "#928374" -bar-bg "#3c3836" -c1 "#fabd2f" -c2 "#b8bb26" -c3 "#fe8019" -c4 "#fb4934"` |
| Everforest | `-bg "#2b3339" -text "#d3c6aa" -subtext "#7a8478" -bar-bg "#374145" -c1 "#a7c080" -c2 "#7fbbb3" -c3 "#dbbc7f" -c4 "#e67e80"` |
| Iceberg | `-bg "#161821" -text "#d2d4de" -subtext "#6b7089" -bar-bg "#1e2132" -c1 "#84a0c6" -c2 "#a093c7" -c3 "#89b8c2" -c4 "#e27878"` |
| Deep Ocean | `-bg "#0f172a" -text "#e2e8f0" -subtext "#64748b" -bar-bg "#1e293b" -c1 "#38bdf8" -c2 "#22c55e" -c3 "#f59e0b" -c4 "#ef4444"` |
| Solarized | `-bg "#002b36" -text "#93a1a1" -subtext "#586e75" -bar-bg "#073642" -c1 "#268bd2" -c2 "#859900" -c3 "#b58900" -c4 "#dc322f"` |
| One Dark | `-bg "#282c34" -text "#abb2bf" -subtext "#5c6370" -bar-bg "#3a3f4b" -c1 "#61afef" -c2 "#98c379" -c3 "#e5c07b" -c4 "#e06c75"` |
| Material | `-bg "#263238" -text "#eeffff" -subtext "#546e7a" -bar-bg "#37474f" -c1 "#82aaff" -c2 "#c3e88d" -c3 "#ffcb6b" -c4 "#f07178"` |

---

## Using the Interactive Web Generator

Not sure which theme to pick? Use the **[Interactive TakaTime Generator](https://rtarun3606k.github.io/TakaTime/)** to preview all themes in real-time and copy the exact command for your workflow.

---

*Generated for [TakaTime](https://github.com/Rtarun3606k/TakaTime) — The Open Source, Self-Hosted WakaTime Alternative.*
