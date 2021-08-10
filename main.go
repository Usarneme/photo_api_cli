package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

type album struct {
	AlbumId      int    `json:"albumId"`
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

// HELPER to read from the json file which is actually just an array of objects
// REMOVE ME TODO
func readLines() ([]string, error) {
	file, err := os.Open("id2.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Validate(inputs []string) (bool, error) {
	if len(inputs) > 1 {
		if inputs[1] == "--help" {
			// check for usage help requested
			return true, errors.New("[usage] - You can call this CLI application with an optional album number. \n[usage] - Either provide a number: go run photo_api_cli.go 8\n[usage] - or, with no arguments: go run photo_api_cli.go\n")
		}
		// check for non-numerical album ID
		_, integerError := strconv.Atoi(os.Args[1])
		if integerError != nil {
			var errString string = "[usage error] - You did not enter a valid number for the album ID. \n[usage error] - Received: " + integerError.Error() + ". \n[usage error] - Please try again or to see more info, try: go run photo_api_cli.go --help\n"
			return false, errors.New(errString)
		}
	}
	// check for too many arguments
	if len(inputs) > 2 {
		return false, errors.New("[usage error] - You can call this CLI application with an optional album number. \n[usage] - Either provide a number: go run photo_api_cli.go 8\n[usage] - or, with no arguments: go run photo_api_cli.go\n")
	}
	// otherwise inputs are valid
	return true, nil
}

func main() {
	color.Yellow("Welcome to the Photos API CLI Tool\n")
	fmt.Println(os.Args)
	// REMOVE ME START TODO
	// lines, readError := readLines()
	// if readError != nil {
	// 	log.Fatalf("readLines: %s", readError)
	// }
	// for i, line := range lines {
	// 	fmt.Println(i, line)
	// }
	// os.Exit(0)
	// REMOVE ME END TODO

	inputsValid, inputsError := Validate(os.Args)
	if inputsError != nil {
		fmt.Print(inputsError)
		if !inputsValid {
			// user called the script with too many, too few, or the wrong arguments
			os.Exit(1)
		} else {
			// user asked for help, not an error condition
			os.Exit(0)
		}
	}

	var albumID int
	var url string

	if len(os.Args) < 2 {
		albumID = 1
		fmt.Printf("Searching for photos from the beginning album.\n")
		url = "https://jsonplaceholder.typicode.com/photos?albumId=1"
	} else {
		fmt.Printf("Searching for photos starting at album ID #%d.\n", albumID)
		url = "https://jsonplaceholder.typicode.com/photos?albumId=" + os.Args[1]
	}

	fmt.Printf("Querying URL: %s\n", url)
	os.Exit(0)

	client := http.Client{
		Timeout: time.Second * 2,
	}
	req, requestError := http.NewRequest(http.MethodGet, url, nil)
	if requestError != nil {
		log.Fatal(requestError)
		os.Exit(1)
	}

	res, getError := client.Do(req)
	if getError != nil {
		log.Fatal(getError)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readError := ioutil.ReadAll(res.Body)
	if readError != nil {
		log.Fatal(readError)
	}

	// body1 := album{}
	// jsonError := json.Unmarshal(body, &body1)
	// if jsonError != nil {
	// 	log.Fatal(jsonError)
	// }

	fmt.Printf("got data: %s \n", body)

	os.Exit(0)
}
