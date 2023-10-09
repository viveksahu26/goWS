package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
)

// list of github URLs
// download all github repo

func main() {
	// List of GitHub repository URLs
	repoURLs := []string{
		"https://github.com/viveksahu26/kubewatchpods",
		"https://github.com/viveksahu26/kuber_explore",
		"https://github.com/viveksahu26/virtual_terminal",
		"https://github.com/viveksahu26/testing_in_go",
		"https://github.com/viveksahu26/url_shortner",
		"https://github.com/kubernetes/kubernetes",
		"https://github.com/sigstore/cosign",
		"https://github.com/viveksahu26/news-app",
	}

	startTime := time.Now()
	fmt.Println("Start time: ", startTime)
	for _, url := range repoURLs {
		fmt.Println("> RepoURL: ", url)
		err := cloneRepo(url)
		if err != nil {
			fmt.Printf("Error cloning %s: %s\n", url, err)
		} else {
			fmt.Printf("Cloned %s\n\n", url)
		}
	}
	endTime := time.Now()
	fmt.Println("End time: ", endTime)
	totalTime := endTime.Sub(startTime)
	fmt.Printf("\n Total time taken: %s\n", totalTime)
	// Total time taken: 10m13.327431315s
}

func cloneRepo(url string) error {
	// Remove "https://github.com/" from the URL to get the repo path
	repoPath := strings.TrimPrefix(url, "https://github.com/")
	fmt.Println("RepoPath: ", repoPath)

	// Create a directory for the repository
	err := os.MkdirAll(repoPath, os.ModePerm)
	if err != nil {
		return err
	}

	// Clone the repository
	_, err = git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	return nil
}
