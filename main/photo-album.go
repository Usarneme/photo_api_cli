package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"

	photo_album "github.com/Usarneme/photo_api_cli"
)

func main() {
	color.Yellow("Welcome to the Photos API CLI Tool\n\n")
	// check if the user asked for usage/syntax help
	if len(os.Args) > 1 {
		if os.Args[1] == "--help" {
			fmt.Print("[usage] - You can call this CLI application with an optional album number. \n[usage] - Either provide a number: go run photo_api_cli.go 8\n[usage] - or, with no arguments: go run photo_api_cli.go\n")
			os.Exit(0)
		}
	}

	// validate that the script was called with correct arguments
	inputsAreValid, inputsError := photo_album.Validate(os.Args)
	if inputsError != nil || !inputsAreValid {
		// user called the script with too many or the wrong type of arguments
		fmt.Print(inputsError)
		os.Exit(1)
	}

	// get the formatted url from the passed-in argument
	var albumID string
	if len(os.Args) == 1 {
		albumID = "1"
	} else {
		albumID = os.Args[1]
	}
	url := photo_album.FormatUrl(albumID)

	// fmt.Printf("Querying URL: %s\n", url)
	result, queryError := photo_album.MakeRequest(url)
	if queryError != nil {
		fmt.Println("Error retrieving from the API. Please try again later.")
		log.Fatal(queryError)
		os.Exit(1)
	}

	photo_album.PrintResults(result)
	os.Exit(0)
}
