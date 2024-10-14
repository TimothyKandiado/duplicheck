package main

import (
	"fmt"
	"os"
)

func groupFilesBySize(path string, filesGroupedBySize map[int64]FileList) {
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

// Groups files based on whether they have the same file size
func groupFilesWithSameSize(filesGroupedBySize map[int64]FileList) []FileList {
	sameSizeFiles := make([]FileList, 0)
	for _, fileList := range filesGroupedBySize {
		if len(fileList) < 2 {
			continue
		}

		sameSizeFiles = append(sameSizeFiles, fileList)
	}

	fmt.Printf("Found %d group of files with same size\n", len(sameSizeFiles))
	return sameSizeFiles
}
