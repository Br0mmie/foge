package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var rootLevelFolders = []string{
	"EPT",
	"IPT",
}

var eptIptLevelFolders = []string{
	"evidence",
	"logs",
	"scans",
	"scope",
	"tools",
}

var evidenceLevelFolders = []string{
	"credentials",
	"data",
	"screenshots",
}

func create_folders(path string, name string) {
	switch name {
	case "":
		for _, name := range rootLevelFolders {
			full := filepath.Join(path, name)
			if err := os.MkdirAll(full, os.ModePerm); err != nil {
				log.Fatalf("Failed to create %s: %v", full, err)
			}
		}
	case "EPT", "IPT":
		parent := filepath.Join(path, name)
		for _, sub := range eptIptLevelFolders {
			full := filepath.Join(parent, sub)
			if err := os.MkdirAll(full, os.ModePerm); err != nil {
				log.Fatalf("Failed to create %s: %v", full, err)
			}
		}

	case "EPT/evidence", "IPT/evidence":
		parent := filepath.Join(path, name[:3])
		for _, sub := range evidenceLevelFolders {
			full := filepath.Join(parent, "evidence", sub)
			if err := os.MkdirAll(full, os.ModePerm); err != nil {
				log.Fatalf("Failed to create %s: %v", full, err)
			}
		}
	}
}

func folder_suite(path string) {
	create_folders(path, "")
	create_folders(path, "EPT")
	create_folders(path, "IPT")
	create_folders(path, "EPT/evidence")
	create_folders(path, "IPT/evidence")
}

func main() {

	if len(os.Args) == 1 {
		fmt.Print("No path specified. Would you like to create it?: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			choice := strings.TrimSpace(scanner.Text())

			switch choice {
			case "y":
				if scanner.Scan() {
					input := strings.TrimSpace(scanner.Text())
					if input == "" {
						fmt.Println("No folder name provided. Exiting.")
						return
					}
					wd, err := os.Getwd()
					if err != nil {
						log.Fatalf("Failed to get current working directory: %v", err)
					}
					path := filepath.Join(wd, input)
					if err := os.MkdirAll(path, os.ModePerm); err != nil {
						log.Fatalf("Failed to create project folder %s: %v", path, err)
					}
					folder_suite(path)
					fmt.Printf("Folder structure created in: %s\n", path)
					return
				}
				if err := scanner.Err(); err != nil {
					log.Fatalf("Error reading input: %v", err)
				}
				return
			case "n":
				wd, _ := os.Getwd()
				folder_suite(wd)
			}
		}
	}
	path := os.Args[1]
	if len(path) == 1 && path == "." {
		wd, _ := os.Getwd()
		folder_suite(wd)
	} else {
		folder_suite(os.Args[1])
	}
}
