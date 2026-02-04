package buildimg

import (
	_ "embed"
	"fmt"
	"image"
	"time"

	"github.com/Rtarun3606k/TakaTime/internal/types"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// Updated to accept 'theme'
func LanguageStatsImg(stats []types.LanguageStat, fontData []byte, theme types.ThemeConfig) (image.Image, error) {
	currettime := time.Now()

	// 1. Setup Canvas (Use theme background)
	dc := setupContext(types.CanvasWidth, types.CanvasHeight, theme)

	// 2. Draw Header
	// FIXED: Now passing 'theme' and size '40.0'
	if err := drawHeader(dc, "TakaTime Language Stats", fontData, theme, 40.0); err != nil {
		return nil, err
	}

	// 3. Draw the List of Languages
	// FIXED: Now passing 'theme'
	if err := drawStatsList(dc, stats, fontData, theme); err != nil {
		return nil, err
	}

	// 4. Draw Footer
	footerText := fmt.Sprintf("Last Updated: %s", currettime.Format("2006-01-02 15:04:05"))
	// FIXED: Now passing 'theme'
	if err := drawFooter(dc, footerText, fontData, theme); err != nil {
		return nil, err
	}

	return dc.Image(), nil
}

// --- Helper Functions ---

// Updated to accept theme
func setupContext(w, h int, theme types.ThemeConfig) *gg.Context {
	dc := gg.NewContext(w, h)
	dc.SetHexColor(theme.BackgroundColor) // Use theme!
	dc.Clear()
	return dc
}

func drawHeader(dc *gg.Context, text string, fontData []byte, theme types.ThemeConfig, size float64) error {
	headerFace, err := loadFontFace(fontData, size)
	if err != nil {
		return err
	}
	dc.SetFontFace(headerFace)
	dc.SetHexColor(theme.TextColor)

	textWidth, _ := dc.MeasureString(text)
	dc.DrawString(text, (float64(dc.Width())-textWidth)/2, size+40)
	return nil
}

// Updated to accept theme
func drawStatsList(dc *gg.Context, stats []types.LanguageStat, fontData []byte, theme types.ThemeConfig) error {
	listFace, err := loadFontFace(fontData, 20)
	if err != nil {
		return fmt.Errorf("failed to load list font: %w", err)
	}
	dc.SetFontFace(listFace)

	startY := 150.0
	gap := 60.0

	for i, stat := range stats {
		y := startY + (float64(i) * gap)
		drawRow(dc, stat, y, theme) // Pass theme down
	}
	return nil
}

// Updated to accept theme
func drawRow(dc *gg.Context, stat types.LanguageStat, y float64, theme types.ThemeConfig) {
	// A. Language Name
	dc.SetHexColor(theme.TextColor)
	dc.DrawString(stat.Name, 50, y)

	// B. Time Text
	dc.SetHexColor(theme.SubTextColor)
	dc.DrawString(stat.Time, 200, y)

	// C. Progress Bar Background
	barX := 350.0
	barW := 350.0
	barH := 15.0
	dc.SetHexColor(theme.BarBackgroundColor)
	dc.DrawRoundedRectangle(barX, y-10, barW, barH, 4)
	dc.Fill()

	// D. Progress Bar Fill
	dc.SetHexColor(stat.Color)
	fillWidth := barW * stat.Percent
	if fillWidth < 10 {
		fillWidth = 10
	}
	dc.DrawRoundedRectangle(barX, y-10, fillWidth, barH, 4)
	dc.Fill()

	// E. Percentage Text
	dc.SetHexColor(theme.TextColor)
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

// Updated to accept theme
func drawFooter(dc *gg.Context, text string, fontData []byte, theme types.ThemeConfig) error {
	footerFace, err := loadFontFace(fontData, 12)
	if err != nil {
		return err
	}
	dc.SetFontFace(footerFace)

	dc.SetHexColor(theme.SubTextColor) // Use theme subtext
	w, h := float64(dc.Width()), float64(dc.Height())
	dc.DrawStringAnchored(text, w-10, h-10, 1.0, 0.0)
	return nil
}

