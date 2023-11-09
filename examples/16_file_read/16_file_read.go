package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("examples/16_file_read/file.json")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("content:")
		fmt.Println(string(content))
	}

	// not-existing file
	content, err = os.ReadFile("examples/16_file_read/unknown.json")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("content:")
		fmt.Println(string(content))
	}
}
