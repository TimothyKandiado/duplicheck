// program for checking duplicate files
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileList []string

var duplicateFiles []FileList

func init() {
	duplicateFiles = make([]FileList, 0)
}

func main() {
	fmt.Println("Starting duplicheck....")

	args := os.Args

	if len(args) > 1 {
		paths := args[1:]

		for _, path := range paths {
			searchDuplicateFiles(path)
		}
	} else {
		path, err := os.Getwd()

		if err != nil {
			fmt.Println(err)
			return
		}

		searchDuplicateFiles(path)
	}

	isolateDuplicateBySizeFiles()
	findDuplicateByHashFiles()
	handleDuplicateFiles()

	fmt.Println("Exiting duplicheck....")
}

func searchDuplicateFiles(path string) {
	directories, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, dir := range directories {
		filepath := fmt.Sprintf("%v/%v", path, dir.Name())
		if !dir.IsDir() {
			groupFilesBySize(filepath)

			continue
		}

		//fmt.Println(dir.Name(), "_ Folder _")
		searchDuplicateFiles(filepath)
	}

}

func handleDuplicateFiles() {
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
		handleDuplicateFiles()
	}

}
