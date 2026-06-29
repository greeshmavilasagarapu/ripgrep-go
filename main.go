package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Current directory
	word := os.Args[1]
	currDir, _ := os.Getwd()
	fmt.Println(currDir)
	// Home directory from environment variable
	home := os.Getenv("HOME")
	fmt.Println(home)
	// Read all files and folders in current directory
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory")
		return
	}
	fmt.Println("\nFiles in current directory:")
	for _, file := range files {
		fmt.Println(file.Name())
		// Print only Go files
		if strings.HasSuffix(file.Name(), ".go") {
			data, _ := os.ReadFile(file.Name())
			if strings.Contains(string(data), word) {
				fmt.Println("Found in:", file.Name())
			}
		}
	}
}
