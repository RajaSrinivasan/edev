package tools

import (
	"fmt"
	"testing"
)

func TestToJSON(t *testing.T) {

	fj, _ := ToJSON("test.go")
	fmt.Printf("%s\n", fj)

	fj, _ = ToJSON("file.go")
	fmt.Printf("%s\n", fj)

}

func TestFromJSON(t *testing.T) {

	fj, _ := ToJSON("file_test.go")
	ofn, _ := FromJSON(fj)
	fmt.Printf("%s\n", ofn)

}