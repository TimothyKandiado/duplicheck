// program for checking duplicate files
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileList []string // List of strings (file paths to be precise)

func main() {
	fmt.Println("Starting duplicheck....")

	args := os.Args

	// Map with file size as key, and slice of file names as value
	filesGroupedBySize := make(map[int64]FileList)

	if len(args) > 1 {
		paths := args[1:]

		for _, path := range paths {
			searchPotentialDuplicateFiles(path, filesGroupedBySize)
		}
	} else {
		path, err := os.Getwd()

		if err != nil {
			fmt.Println(err)
			return
		}

		searchPotentialDuplicateFiles(path, filesGroupedBySize)
	}

	sameSizeFiles := groupFilesWithSameSize(filesGroupedBySize)

	duplicateFiles := findDuplicateFilesByHash(sameSizeFiles)

	handleDuplicateFiles(duplicateFiles)

	fmt.Println("Exiting duplicheck....")
}

func searchPotentialDuplicateFiles(path string, filesGroupedBySize map[int64]FileList) {
	directories, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, dir := range directories {
		filepath := fmt.Sprintf("%v/%v", path, dir.Name())
		if !dir.IsDir() {
			groupFilesBySize(filepath, filesGroupedBySize)

			continue
		}

		//fmt.Println(dir.Name(), "_ Folder _")
		searchPotentialDuplicateFiles(filepath, filesGroupedBySize)
	}
}

func handleDuplicateFiles(duplicateFiles []FileList) {
	if len(duplicateFiles) == 0 {
		fmt.Println("No duplicates were found")
		return
	}
	fmt.Printf("%d file duplicates were found\n", len(duplicateFiles))
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Choose handling mode\n0 for Auto\n1 for Manual")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	input = strings.Replace(input, "\n", "", -1)
	input = strings.TrimSpace(input)
	//fmt.Println("input: ", input)

	if input == "0" || input == "auto" {
		autoDuplicateHandler(duplicateFiles)
	} else if input == "1" || input == "manual" {
		manualDuplicateHandler(duplicateFiles)
	} else {
		fmt.Println("Invalid Selection")
		handleDuplicateFiles(duplicateFiles)
	}

}
