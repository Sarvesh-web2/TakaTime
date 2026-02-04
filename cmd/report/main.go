package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
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

	theme := types.DefaultTheme()

	// Flags
	flag.StringVar(&theme.BackgroundColor, "bg", theme.BackgroundColor, "Background Color")
	flag.StringVar(&theme.Color1, "c1", theme.Color1, "Primary Color")
	flag.Parse()

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
			return buildimg.DrawListCard("Language Stats", langs, fontData, time.Now(), theme)
		})

		// Job 2: Projects
		handleImageJob("Projects", "public/taka-projects.png", gistToken, targetRepo, func() (image.Image, error) {
			return buildimg.DrawListCard("Top Projects", projects, fontData, time.Now(), theme)
		})

		// Job 3: Time Grid (2x2 View)
		handleImageJob("Time Stats", "public/taka-time.png", gistToken, targetRepo, func() (image.Image, error) {
			return buildimg.DrawTimeCard(timeStats, fontData, time.Now(), theme)
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
	SaveImage(name+".png",img)

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

func SaveImage(filename string, img image.Image) error {
	// 1. Create the file
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	// 2. Encode the image as PNG
	if err := png.Encode(f, img); err != nil {
		return fmt.Errorf("failed to encode PNG: %w", err)
	}

	fmt.Printf("✅ Saved debug image: %s\n", filename)
	return nil
}
