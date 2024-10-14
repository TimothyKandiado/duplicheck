package main

import (
	"crypto/sha1"
	"fmt"
	"os"
)

type HashCode [20]byte

// goes through the 2D array of files with similar sizes,
// calculates their hash, then groups the files with the same hash
func findDuplicateFilesByHash(sameSizeFiles []FileList) []FileList {
	duplicatesByHash := make([]FileList, 0)

	for _, files := range sameSizeFiles {
		hashDuplicateWithSameSize := make(map[HashCode]FileList)

		// group files by their hash value
		for _, filepath := range files {
			hash, err := hashFile(filepath)

			if err != nil {
				fmt.Println(err)
				continue
			}

			files, exist := hashDuplicateWithSameSize[hash]
			if exist {
				files = append(files, filepath)
				hashDuplicateWithSameSize[hash] = files
			} else {
				hashDuplicateWithSameSize[hash] = []string{filepath}
			}
		}

		// isolate duplicate files
		for _, files := range hashDuplicateWithSameSize {
			if len(files) < 2 {
				continue
			}

			duplicatesByHash = append(duplicatesByHash, files)
		}
	}

	fmt.Printf("Found %d group of files with same hash\n", len(duplicatesByHash))

	return duplicatesByHash
}

func hashFile(path string) (HashCode, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return [20]byte{}, err
	}

	hash := sha1.Sum(data)

	return hash, nil
}
