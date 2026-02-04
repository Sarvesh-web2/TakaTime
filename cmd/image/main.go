package main

import (
	"fmt"
	"image/color"
	"log"

	_ "embed"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype" // <--- 2. NEW IMPORT
	"golang.org/x/image/font"
)

type LanguageStat struct {
	Name    string
	Time    string
	Hours   float64 // Used for bar calculation
	Percent float64 // Used for bar width (0.0 to 1.0)
	Color   string  // Hex color for the bar
}

// //go:embed font.ttf

//go:embed FiraCodeNerdFontPropo-Retina.ttf
var fontData []byte

func main() {
	// 1. Setup Canvas (Width x Height)
	const W, H = 800, 500
	dc := gg.NewContext(W, H)

	// 2. Background (GitHub Dark Dimmed: #0d1117)
	dc.SetHexColor("#0d1117")
	dc.Clear()

	// 3. Load Font (MUST exist in the folder)
	// You can try "Arial.ttf" or "JetBrainsMono.ttf"
	// if err := dc.LoadFontFace("font.ttf", 40); err != nil {
	// if err := dc.LoadFontFace("FiraCodeNerdFontPropo-Retina.ttf", 40); err != nil {
	// 	log.Fatal("❌ Error: Could not load 'font.ttf'. Please place a .ttf file in this folder.\n", err)
	// }

	headerFace, err := loadFontFace(fontData, 40)
	if err != nil {
		log.Fatal("some problem loading font!", err)
	}
	dc.SetFontFace(headerFace)

	// 4. Draw Header (Centered)
	dc.SetColor(color.White)
	header := "TakaTime Language Stats"
	textWidth, _ := dc.MeasureString(header)
	dc.DrawString(header, (W-textWidth)/2, 60) // Center horizontally

	// 5. Dummy Data (Top 5)
	stats := []LanguageStat{
		{"Go", "14h 32m", 14.5, 0.75, "#00ADD8"},        // Cyan
		{"Rust", "8h 15m", 8.25, 0.56, "#dea584"},       // Orange/Rust
		{"TypeScript", "5h 45m", 5.75, 0.39, "#3178c6"}, // Blue
		{"Lua", "3h 20m", 3.33, 0.23, "#000080"},        // Navy
		{"Markdown", "1h 10m", 1.16, 0.08, "#ffffff"},   // White
		{"Others", "1h 10m", 1.16, 0.25, "#ffffc6"},     // White
	}

	// 6. Draw Loop
	startY := 150.0
	gap := 60.0
	// dc.LoadFontFace("FiraCodeNerdFontPropo-Retina.ttf", 20) // Smaller font for list
	listFace, _ := loadFontFace(fontData, 20)
	dc.SetFontFace(listFace)
	for i, stat := range stats {
		y := startY + (float64(i) * gap)
		drawRow(dc, stat, y)
	}

	// 7. Save
	if err := dc.SavePNG("test-dashboard.png"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ Generated 'test-dashboard.png' successfully!")
}

func drawRow(dc *gg.Context, stat LanguageStat, y float64) {
	// A. Language Name (Left aligned)
	dc.SetColor(color.White)
	dc.DrawString(stat.Name, 50, y)

	// B. Time Text (Aligned next to name or fixed column)
	dc.SetHexColor("#8b949e") // Gray text
	dc.DrawString(stat.Time, 200, y)

	// C. Progress Bar Background
	barX := 350.0
	barW := 350.0
	barH := 15.0
	dc.SetHexColor("#30363d") // Dark Gray container
	dc.DrawRoundedRectangle(barX, y-10, barW, barH, 4)
	dc.Fill()

	// D. Progress Bar Fill (Actual Usage)
	dc.SetHexColor(stat.Color)
	fillWidth := barW * stat.Percent
	if fillWidth < 10 {
		fillWidth = 10
	} // Min width visibility
	dc.DrawRoundedRectangle(barX, y-10, fillWidth, barH, 4)
	dc.Fill()

	// E. Percentage Text (Right of bar)
	dc.SetColor(color.White)
	percentStr := fmt.Sprintf("%.1f%%", stat.Percent*100)
	dc.DrawString(percentStr, barX+barW+20, y)
}

func loadFontFace(fontBytes []byte, points float64) (font.Face, error) {
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size:    points,
		Hinting: font.HintingFull,
	})
	return face, nil
}

