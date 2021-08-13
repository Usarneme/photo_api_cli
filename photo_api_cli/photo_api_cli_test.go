package photo_api_cli_test

import (
	"errors"
	photo_album "github.com/Usarneme/photo_api_cli"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Validate, FormatUrl, MakeRequest, PrintResults
func TestValidateWithNoArgs(t *testing.T) {
	// args[0] is the go run command itself
	args := []string{"/var/folders/1/2/3/T/go-build/b/exe/main"}
	argsAreValid, validationError := photo_album.Validate(args)
	assert.Equal(t, nil, validationError)
	assert.Equal(t, true, argsAreValid)
}

func TestValidateWithAlbumIdEqualToOne(t *testing.T) {
	args := []string{"/var/folders/1/2/3/T/go-build/b/exe/main", "1"}
	argsAreValid, validationError := photo_album.Validate(args)
	assert.Equal(t, nil, validationError)
	assert.Equal(t, true, argsAreValid)
}

func TestValidateWithTooManyArgs(t *testing.T) {
	args := []string{"/var/folders/1/2/3/T/go-build/b/exe/main", "1", "2", "3"}
	argsAreValid, validationError := photo_album.Validate(args)
	errorMessage := errors.New("[usage error] - You can call this CLI application with an optional album number. \n[usage] - Either provide a number: go run photo_api_cli.go 8\n[usage] - or, with no arguments: go run photo_api_cli.go\n")
	assert.Equal(t, errorMessage, validationError)
	assert.Equal(t, false, argsAreValid)
}

func TestValidateWithNonIntegerAlbumID(t *testing.T) {
	args := []string{"/var/folders/1/2/3/T/go-build/b/exe/main", "blah"}
	argsAreValid, validationError := photo_album.Validate(args)
	errorMessage := errors.New("[usage error] - You did not enter a valid number for the album ID.\n[usage error] - Please try again or to see more info, try: go run photo_api_cli.go --help\n")
	assert.Equal(t, errorMessage, validationError)
	assert.Equal(t, false, argsAreValid)
}

func TestFormatUrlWithNoAlbumID(t *testing.T) {
	url := photo_album.FormatUrl("")
	assert.Equal(t, "https://jsonplaceholder.typicode.com/photos?albumId=1", url)
}

func TestFormatUrlWithAlbumIDEqualToFour(t *testing.T) {
	url := photo_album.FormatUrl("4")
	assert.Equal(t, "https://jsonplaceholder.typicode.com/photos?albumId=4", url)
}
