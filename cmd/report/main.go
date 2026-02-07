package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"image"
	"log"
	"os"
	"strings"
	"time"

	buildimg "github.com/Rtarun3606k/TakaTime/internal/BuildImg"
	dbqueryv2 "github.com/Rtarun3606k/TakaTime/internal/DBQueryV2"
	gogist "github.com/Rtarun3606k/TakaTime/internal/GoGist"
	utils "github.com/Rtarun3606k/TakaTime/internal/Utils"
	"github.com/Rtarun3606k/TakaTime/internal/db"
	"github.com/Rtarun3606k/TakaTime/internal/types"
)

//go:embed FiraCodeNerdFontPropo-Retina.ttf
var fontData []byte

func main() {

	var themeFlag string
	var customBg, customText, customSubText, customBarBg string
	var customC1, customC2, customC3, customC4 string

	//----------flags
	days := flag.Int("days", 0, "Number of past days to include (0 = today)")
	flag.StringVar(&themeFlag, "theme", "dark", "Base theme: dark, light, dracula, nord, gruvbox, monokai, cyberpunk")

	// Custom Overrides
	flag.StringVar(&customBg, "bg", "", "Override Background Color")
	flag.StringVar(&customText, "text", "", "Override Text Color")
	flag.StringVar(&customSubText, "subtext", "", "Override Sub-Text Color")
	flag.StringVar(&customBarBg, "bar-bg", "", "Override Bar Background Color")

	flag.StringVar(&customC1, "c1", "", "Override Color 1 (Primary)")
	flag.StringVar(&customC2, "c2", "", "Override Color 2")
	flag.StringVar(&customC3, "c3", "", "Override Color 3")
	flag.StringVar(&customC4, "c4", "", "Override Color 4")

	flag.Parse() //  Parse flags first!
	log.Println("days as a flag is discontinued please stop adding it as a flag. remove it from your work flow file check out readme https://github.com/Rtarun3606k/Takatime for more information ", days)

	// 2. Initialize the Final 'theme' Variable
	var theme types.ThemeConfig

	// Decide the base: Start with a full "Light" or "Dark" template
	switch themeFlag {
	case "light":
		theme = types.LightTheme
	case "dracula":
		theme = types.DraculaTheme
	case "nord":
		theme = types.NordTheme
	case "gruvbox":
		theme = types.GruvboxTheme
	case "monokai":
		theme = types.MonokaiTheme
	case "cyberpunk":
		theme = types.CyberpunkTheme
	default:
		// Default to Dark if unknown or explicitly "dark"
		theme = types.DefaultTheme()
	}

	// Apply Overrides
	if customBg != "" {
		theme.BackgroundColor = customBg
	}

	// 3. Apply Overrides
	if customBg != "" {
		theme.BackgroundColor = customBg
	}
	if customText != "" {
		theme.TextColor = customText
	}
	if customSubText != "" {
		theme.SubTextColor = customSubText
	}
	if customBarBg != "" {
		theme.BarBackgroundColor = customBarBg
	}

	if customC1 != "" {
		theme.Color1 = customC1
	}
	if customC2 != "" {
		theme.Color2 = customC2
	}
	if customC3 != "" {
		theme.Color3 = customC3
	}
	if customC4 != "" {
		theme.Color4 = customC4
	}

	// Connect
	mongoURI := os.Getenv("MONGO_URI")
	// gistID := os.Getenv("GIST_ID")
	gistToken := os.Getenv("GIST_TOKEN")
	targetRepo := os.Getenv("TARGET_REPO")

	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is required")
	}

	client, err := db.ConnectToDataBase(mongoURI)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	if gistToken != "" && targetRepo != "" {
		fmt.Printf("Updating README for %s...\n", targetRepo)

		if err := dbqueryv2.RunMigrations(client); err != nil {
			log.Printf("Migration warning: %v", err)
			// Don't crash, just log it. The report can still run.
		}

		name := strings.Split(targetRepo, "/")

		// 1. Fetch Real Data
		projects, err := dbqueryv2.GetListStats(client, "project", 5, theme)
		if err != nil {
			log.Println("Proj Error:", err)
		}

		langs, err := dbqueryv2.GetListStats(client, "language", 5, theme)
		if err != nil {
			log.Println("Lang Error:", err)
		}

		editors, err := dbqueryv2.GetListStats(client, "editor", 3, theme)
		if err != nil {
			log.Println("Editor Error:", err)
		}

		osStats, err := dbqueryv2.GetListStats(client, "os", 3, theme)
		if err != nil {
			log.Println("OS Error:", err)
		}

		timeStats, err := dbqueryv2.GetTimeStats(client)
		if err != nil {
			log.Println("Time Error:", err)
		}

		// job 1 : Languages
		handleImageJob("Languages", "public/taka-languages.png", gistToken, targetRepo, func() (image.Image, error) {
			return buildimg.DrawListCard("Language Stats", langs, fontData, time.Now(), theme, false)
		})

		// Job 2: Projects
		handleImageJob("Projects", "public/taka-projects.png", gistToken, targetRepo, func() (image.Image, error) {
			return buildimg.DrawListCard("Top Projects", projects, fontData, time.Now(), theme, true)
		})

		// Job 3: Time Grid (2x2 View)
		handleImageJob("Time Stats", "public/taka-time.png", gistToken, targetRepo, func() (image.Image, error) {
			return buildimg.DrawTimeCard(timeStats, fontData, time.Now(), theme, name[0])
		})

		// Job 4: Tech Stack (Editors Left / OS Right)
		handleImageJob("Tech Stack", "public/taka-tech.png", gistToken, targetRepo, func() (image.Image, error) {
			// Pass both lists to the split-view generator
			return buildimg.DrawTechCard(editors, osStats, fontData, time.Now(), theme)
		})

		content := utils.GenerateOutput()

		errr := gogist.UpdateReadMe(gistToken, targetRepo, content)
		if errr != nil {
			fmt.Println("Some error occured while updating readme ", err)
		}
		fmt.Println("README Updated Successfully!")
	} else {
		fmt.Println("Skipping README update (GIST_TOKEN or TARGET_REPO missing)")
	}

	//
	// updateContentError := gogist.UpdateGist(gistToken, gistID, content)
	// if updateContentError != nil {
	// 	log.Fatalln("Some error occured gogist", updateContentError)
	// 	return
	// }
	//

}

func handleImageJob(name, path, token, repo string, generator func() (image.Image, error)) {
	fmt.Printf("Processing %s...\n", name)

	// 1. Generate Image
	img, err := generator()
	if err != nil {
		log.Printf("Gen Error (%s): %v\n", name, err)
		return
	}
	// SaveImage(name+".png", img)

	// 2. Format Config (Using your utils package)
	cfg, err := utils.FormmatUpload(token, repo, path, "main", "Update "+name)
	if err != nil {
		log.Printf("Config Error (%s): %v\n", name, err)
		return
	}

	// 3. Upload with FRESH Timeout (Critical for loops!)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel() // Cancels this specific context when function exits

	if err := gogist.UploadImageToGitHub(ctx, img, cfg); err != nil {
		log.Printf("Upload Error (%s): %v\n", name, err)
	} else {
		fmt.Printf("Uploaded: %s\n", path)
	}
}

// func SaveImage(filename string, img image.Image) error {
// 	// 1. Create the file
// 	f, err := os.Create(filename)
// 	if err != nil {
// 		return fmt.Errorf("failed to create file: %w", err)
// 	}
// 	defer f.Close()
//
// 	// 2. Encode the image as PNG
// 	if err := png.Encode(f, img); err != nil {
// 		return fmt.Errorf("failed to encode PNG: %w", err)
// 	}
//
// 	fmt.Printf("Saved debug image: %s\n", filename)
// 	return nil
// }
