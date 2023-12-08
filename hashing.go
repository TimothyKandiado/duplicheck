package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

type HashCode [16]byte

var duplicatesByHash []FileList

func init() {
	duplicatesByHash = make([]FileList, 0)
}

func findDuplicateByHashFiles() {
	for _, files := range duplicatesBySize {
		hashDuplicateWithSameSize := make(map[HashCode]FileList)

		// group hash duplicates with same size
		for _, filepath := range files {
			hash, err := hashFile(filepath)

			if err != nil {
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

		// isolate and store files with duplicates
		for _, files := range hashDuplicateWithSameSize {
			if len(files) < 2 {
				continue
			}

			duplicatesByHash = append(duplicatesByHash, files)
		}
	}
	fmt.Printf("Found %d group of files with same hash\n", len(duplicatesByHash))
	duplicateFiles = duplicatesByHash
}

func hashFile(path string) ([16]byte, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return [16]byte{}, err
	}

	hash := md5.Sum(data)

	return hash, nil
}
