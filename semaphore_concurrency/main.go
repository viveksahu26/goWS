package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
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
		//"https://github.com/viveksahu26/news-app",
	}

	startTime := time.Now()
	fmt.Println("Start time: ", startTime)

	ctx := context.Background()

	g, ctx := errgroup.WithContext(ctx)
	sem := semaphore.NewWeighted(int64(runtime.NumCPU()))
	fmt.Println("Number of CPU Core: ", runtime.NumCPU())

	for _, url := range repoURLs {
		url := url
		g.Go(func() error {
			if err := sem.Acquire(ctx, 1); err != nil {
				return err
			}
			defer sem.Release(1)

			fmt.Println("> RepoURL: ", url)
			err := CloneRepo(url)
			if err != nil {
				fmt.Printf("Error cloning %s: %s\n", url, err)
			} else {
				fmt.Printf("Cloned %s\n\n", url)
			}

			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("err: ", err)
	}

	endTime := time.Now()
	fmt.Println("End time: ", endTime)
	totalTime := endTime.Sub(startTime)
	fmt.Printf("\n Total time taken: %s\n", totalTime)
	// Total time taken: 10m13.327431315s
	// Total time taken: 7m34.270479508s
}

func CloneRepo(url string) error {
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
