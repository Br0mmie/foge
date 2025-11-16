package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	path := os.Args[1]

	if len(path) == 1 && path == "." {
		wd, _ := os.Getwd()
		folder_suite(wd)
	} else {
		fmt.Println(path)
	}
}
