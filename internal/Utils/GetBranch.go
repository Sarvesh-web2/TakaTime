package utils

import (
	"os/exec"
	"strings"
)

func GetGitBranch(dir string) (string, error) {

	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	branch := strings.TrimSpace(string(output))
	return branch, nil
}
