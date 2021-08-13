package photo_api_cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Album struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

func Validate(inputs []string) (bool, error) {
	// check for too many arguments
	if len(inputs) > 2 {
		return false, errors.New("[usage error] - You can call this CLI application with an optional album number. \n[usage] - Either provide a number: go run photo_api_cli.go 8\n[usage] - or, with no arguments: go run photo_api_cli.go\n")
	}
	if len(inputs) > 1 {
		// check for a non-numerical album ID
		_, integerError := strconv.Atoi(inputs[1])
		if integerError != nil {
			var errString string = "[usage error] - You did not enter a valid number for the album ID.\n[usage error] - Please try again or to see more info, try: go run photo_api_cli.go --help\n"
			return false, errors.New(errString)
		}
	}
	// otherwise inputs are valid
	return true, nil
}

func FormatUrl(albumID string) string {
	if albumID == "" {
		url := "https://jsonplaceholder.typicode.com/photos?albumId=1"
		return url
	}
	url := "https://jsonplaceholder.typicode.com/photos?albumId=" + albumID
	return url
}

func MakeRequest(url string) ([]Album, error) {
	albums := []Album{}

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, requestError := http.NewRequest(http.MethodGet, url, nil)
	if requestError != nil {
		return albums, requestError
	}

	res, getError := client.Do(req)
	if getError != nil {
		return albums, getError
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readError := ioutil.ReadAll(res.Body)
	if readError != nil {
		return albums, readError
	}

	if len(body) == 2 {
		// byte array is empty, no results
		return albums, errors.New("Empty array returned. No data available at that album ID.")
	}

	parseError := json.Unmarshal(body, &albums)
	if parseError != nil {
		return albums, errors.New("Error parsing JSON into array of albums.")
	}

	return albums, nil
}

func PrintResults(rawResults []Album) {
	for _, s := range rawResults {
		fmt.Printf("[%d] %s\n", s.ID, s.Title)
	}
}
