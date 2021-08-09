package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	color.Yellow("Welcome to the Photos API CLI Tool\n")

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		fmt.Printf("[usage] - You can call this CLI application with an optional album number. \n[usage] - Either provide a number: go run photo_api_cli.go 8\n[usage] - or, with no arguments: go run photo_api_cli.go\n")
		os.Exit(0)
	}

	var albumID int
	var err error
	if len(os.Args) < 2 {
		albumID = 0
		fmt.Printf("Searching for photos from the beginning album.\n")
	} else {
		albumID, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("[usage error] - You did not enter a valid number for the album ID. \n[usage error] - Received: %e. \n[usage error] - Please try again or to see more info, try: go run photo_api_cli.go --help\n", err)
			os.Exit(1)
		}
		fmt.Printf("Searching for photos starting at album ID #%d.\n", albumID)
	}

	os.Exit(0)
}
