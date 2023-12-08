// program for checking duplicate files
package main

import (
	"bufio"
	"fmt"
	"hash/adler32"
	"os"
	"strings"
)

// type hashCode uint32
type fileList []string

var files map[uint32]fileList
var duplicateFiles []fileList

func main() {
	fmt.Println("Starting duplicheck....")
	files = make(map[uint32]fileList)

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

	isolateDuplicateFiles()
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
			hash, err := hashFile(filepath)

			if err != nil {
				fmt.Println(err)
				continue
			}

			value, exists := files[hash]

			if exists {
				// fmt.Printf("duplicate found: %v\n", filepath)
				value = append(value, filepath) // add duplicate file to hash code
				files[hash] = value
			} else {
				files[hash] = []string{filepath}
			}

			//fmt.Println(dir.Name(), " _ File _ Hash: ", hash)
			continue
		}

		//fmt.Println(dir.Name(), "_ Folder _")
		searchDuplicateFiles(filepath)
	}

}

func hashFile(path string) (uint32, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return 0, err
	}

	hash := adler32.Checksum(data)

	return hash, nil
}

func isolateDuplicateFiles() {
	duplicateFiles = make([]fileList, 0)
	for _, list := range files {
		if len(list) > 1 {
			//fmt.Println(list)
			duplicateFiles = append(duplicateFiles, list)
		}
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
