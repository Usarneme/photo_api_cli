package main

import (
	"testing"
)

func TestValidateCallingWithNoArgs(t *testing.T) {
	var args []string
	args[0] = "/var/folders/1/2/3/T/go-build/b/exe/main"
	argsAreValid, validationError := Validate(args)
	if !argsAreValid || validationError != nil {
		t.Fatalf(`Error in calling the script with no arguments`)
	}
}
