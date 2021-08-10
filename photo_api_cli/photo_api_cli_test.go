package photo_api_cli_test

import (
	photo_album "github.com/Usarneme/photo_api_cli"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateCallingWithNoArgs(t *testing.T) {
	var args []string
	args[0] = "/var/folders/1/2/3/T/go-build/b/exe/main"
	argsAreValid, validationError := photo_album.Validate(args)
	if !argsAreValid || validationError != nil {
		t.Fatalf(`Error in calling the script with no arguments`)
	}
}
