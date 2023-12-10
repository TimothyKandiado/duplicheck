package main

import (
	"fmt"
	"os"
)

var duplicatesBySize []FileList
var filesGroupedBySize map[int64]FileList

func init() {
	filesGroupedBySize = make(map[int64]FileList)
	duplicatesBySize = make([]FileList, 0)
}

func groupFileBySize(path string) {
	fileInfo, err := os.Stat(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	fileSize := fileInfo.Size()

	fileList, exists := filesGroupedBySize[fileSize]

	if exists {
		fileList = append(fileList, path)
		filesGroupedBySize[fileSize] = fileList
	} else {
		filesGroupedBySize[fileSize] = []string{path}
	}
}

func isolateDuplicateBySizeFiles() {
	for _, fileList := range filesGroupedBySize {
		if len(fileList) < 2 {
			continue
		}

		duplicatesBySize = append(duplicatesBySize, fileList)
	}

	fmt.Printf("Found %d group of files with same size\n", len(duplicatesBySize))
}
