package main

import (
	"fmt"
	"os"
	"time"
)

func keepRecentFile(duplicates FileList) {
	var recentTime time.Time = time.Date(2000, time.January, 1, 1, 1, 1, 1, time.Local)
	var recentIndex int

	for index, filePath := range duplicates {
		fileInfo, err := os.Stat(filePath)

		if err != nil {
			fmt.Println(err)
			return
		}

		time := fileInfo.ModTime()

		if time.Compare(recentTime) >= 0 {
			recentTime = time
			recentIndex = index
		}
	}

	deleteDuplicates(duplicates, recentIndex)
}

func keepShortestPath(duplicates FileList) {
	var shortestIndex int = 0
	var shortestLength int = 1000000

	for index, filePath := range duplicates {
		if len(filePath) < shortestLength {
			shortestIndex = index
			shortestLength = len(filePath)
		}
	}

	deleteDuplicates(duplicates, shortestIndex)
}
