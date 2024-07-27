package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yourusername/zetta"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: zetta git repo create --repo <name> --type <type> --service <url>")
		os.Exit(1)
	}

	serviceURL := flag.String("service", "", "Service URL for repository creation")
	repoName := flag.String("repo", "", "Repository name")
	repoType := flag.String("type", "", "Repository type")

	// Parse command-line flags
	flag.CommandLine.Parse(os.Args[2:])

	if *serviceURL == "" || *repoName == "" {
		fmt.Println("Service URL and repository name are required")
		os.Exit(1)
	}

	sdk := zetta.NewSDK(*serviceURL)
	if err := sdk.CreateRepository(*repoName, *repoType); err != nil {
		log.Fatalf("Error creating repository: %v", err)
	}
	fmt.Println("Repository created and LFS configured successfully.")
}
