package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AutoMode int

const (
	Recent AutoMode = iota
	ShortestPath
)

func autoDuplicateHandler(duplicateFiles []fileList) {
	fmt.Println("Automatic mode")
	fmt.Println("Choose automatic settings")
	fmt.Println("0 to keep most recent file\n1 to keep shortest file path")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	input = strings.Replace(input, "\n", "", -1)
	input = strings.TrimSpace(input)

	var autoMode AutoMode

	switch input {
	case "0":
		autoMode = Recent
	case "1":
		autoMode = ShortestPath
	default:
		fmt.Println("Invalid selection")
		autoDuplicateHandler(duplicateFiles)
		return
	}

	for _, duplicates := range duplicateFiles {
		switch autoMode {
		case Recent:
			keepRecentFile(duplicates)
		case ShortestPath:
			keepShortestPath(duplicates)
		}
	}
}

func manualDuplicateHandler(duplicateFiles []fileList) {
	for _, duplicates := range duplicateFiles {
		deleteDuplicateManually(duplicates)
	}
}

func deleteDuplicateManually(duplicates fileList) {
	fmt.Println("\nSelect file to keep")

	for index, filePath := range duplicates {
		fmt.Printf("[%d] -> %v\n", index, filePath)
	}

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	input = strings.Replace(input, "\n", "", -1)
	input = strings.TrimSpace(input)

	selection, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	if selection < 0 || selection >= len(duplicates) {
		fmt.Println("Invalid input")
		return
	}

	deleteDuplicates(duplicates, selection)
}

func deleteDuplicates(duplicates fileList, excludeIndex int) {
	for index, filePath := range duplicates {
		if index == excludeIndex {
			fmt.Println("Keeping: ", filePath)
			continue
		}

		fmt.Println("Deleting: ", filePath)
		deleteFile(filePath)
	}
}

func deleteFile(filePath string) {
	err := os.Remove(filePath)

	if err != nil {
		fmt.Println(err)
	}
}
