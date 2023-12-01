package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/github"
)

func getPublicRepositories(username string) ([]string, error) {
	ctx := context.Background()

	client := github.NewClient(nil)

	opt := &github.RepositoryListOptions{
		Type: "public",
	}

	var repositories []string

	for {
		repos, resp, err := client.Repositories.List(ctx, username, opt)
		if err != nil {
			return nil, err
		}
		for _, repo := range repos {
			repositories = append(repositories, *repo.SSHURL)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return repositories, nil
}

// ... (rest of the code remains the same)

func main() {
	// GitHub username and access token
	username := "viveksahu26"

	// Get list of repositories
	repositories, err := getPublicRepositories(username)
	if err != nil {
		fmt.Printf("Error getting repositories: %s\n", err)
		return
	}
	fmt.Println("REPOSITORIES: ", repositories)
	// os.Exit(0)

	startTime := time.Now()

	// Clone each repository
	for _, repoURL := range repositories {
		err := CloneRepo(repoURL)
		if err != nil {
			fmt.Printf("Error cloning %s: %s\n", repoURL, err)
		} else {
			fmt.Printf("Cloned %s\n", repoURL)
		}
	}

	endTime := time.Now()
	totalTime := endTime.Sub(startTime)

	fmt.Printf("Total time taken: %s\n", totalTime)
}

func CloneRepo(url string) error {
	urlx := strings.Replace(url, "git@github.com:", "https://github.com/", 1)
	fmt.Println("URLX: ", urlx)
	// Remove "https://github.com/" from the URL to get the repo path
	repoPath := strings.TrimPrefix(urlx, "https://github.com/")
	// Remove "https://github.com/" from the URL to get the repo path
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
