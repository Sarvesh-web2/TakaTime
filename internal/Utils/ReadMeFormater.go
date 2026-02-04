package utils

import "strings"

func GenerateOutput() string {
	var sb strings.Builder

	// 1. Header
	sb.WriteString("# TakaTime Weekly Report\n\n")
	sb.WriteString("Check out my coding activity over the last week!\n\n")

	// 2. The Dashboard (HTML Grid for alignment)
	// We use 'align="center"' to center everything nicely.

	sb.WriteString("<div align=\"center\">\n\n")

	// --- ROW 1: TIME STATS (The Summary) ---
	// Full width image
	sb.WriteString("\n")
	sb.WriteString("<img src=\"./public/taka-time.png\" width=\"100%\" alt=\"Time Stats\" /><br/><br/>\n\n")

	// --- ROW 2: LANGUAGES & PROJECTS (Side by Side) ---
	// We use width="48%" to make them sit next to each other
	sb.WriteString("\n")
	sb.WriteString("<img src=\"./public/taka-languages.png\" width=\"48%\" alt=\"Languages\" />")
	sb.WriteString("&nbsp;&nbsp;") // Small spacer
	sb.WriteString("<img src=\"./public/taka-projects.png\" width=\"48%\" alt=\"Projects\" /><br/><br/>\n\n")

	// --- ROW 3: TECH STACK (Editors & OS) ---
	// Full width image (since it's a split view internally)
	sb.WriteString("\n")
	sb.WriteString("<img src=\"./public/taka-tech.png\" width=\"100%\" alt=\"Tech Stack\" />\n\n")

	sb.WriteString("</div>\n\n")

	// 3. Footer
	sb.WriteString("_Generated automatically by [TakaTime](https://github.com/Rtarun3606k/TakaTime)_")

	return sb.String()
}
