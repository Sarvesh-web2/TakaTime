package gogist

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/png"
	"time"

	"github.com/Rtarun3606k/TakaTime/internal/types"
	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

func UploadImageToGitHub(ctx context.Context, img image.Image, cfg types.UploadStruct) error {
	// 1. Authenticate
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: cfg.Token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// 2. Encode Image to Bytes (In-Memory)
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return fmt.Errorf("failed to encode PNG: %w", err)
	}
	fileContent := buf.Bytes()

	// 3. Check if file exists (We need the SHA to update it)
	// We ignore the error here because if it fails, it usually just means the file doesn't exist yet.
	existingFile, _, _, _ := client.Repositories.GetContents(ctx, cfg.Owner, cfg.Repo, cfg.Path, &github.RepositoryContentGetOptions{Ref: cfg.Branch})

	// 4. Prepare Commit Options
	now := time.Now().Format("2006-01-02")
	msg := fmt.Sprintf("%s [%s]", cfg.CommitMsg, now)

	opts := &github.RepositoryContentFileOptions{
		Message: &msg,
		Content: fileContent,
		Branch:  &cfg.Branch,
		Committer: &github.CommitAuthor{
			Name:  github.String("TakaTime Bot"),
			Email: github.String("bot@takatime.dev"),
		},
	}

	// 5. Execute (Create or Update)
	if existingFile != nil {
		// UPDATE
		opts.SHA = existingFile.SHA
		_, _, err := client.Repositories.UpdateFile(ctx, cfg.Owner, cfg.Repo, cfg.Path, opts)
		if err != nil {
			return fmt.Errorf("update failed: %w", err)
		}
		fmt.Println("🚀 Image updated successfully on GitHub!")
	} else {
		// CREATE
		_, _, err := client.Repositories.CreateFile(ctx, cfg.Owner, cfg.Repo, cfg.Path, opts)
		if err != nil {
			return fmt.Errorf("creation failed: %w", err)
		}
		fmt.Println("✨ Image created successfully on GitHub!")
	}

	return nil
}
