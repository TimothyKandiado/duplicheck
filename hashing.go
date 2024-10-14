package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	"runtime"
)

const hashCodeLength = 20
const channelBuffer = 100

type HashCode [hashCodeLength]byte

// goes through the 2D array of files with similar sizes,
// calculates their hash, then groups the files with the same hash
func findDuplicateFilesByHash(sameSizeFiles []FileList) []FileList {
	duplicatesByHash := make([]FileList, 0)

	filepathChannel := make(chan string, channelBuffer)
	computedHashChannel := make(chan struct {
		HashCode
		string
	}, channelBuffer)

	defer close(filepathChannel)
	defer close(computedHashChannel)

	nullHash := HashCode(make([]byte, hashCodeLength))

	// lambda function for computing hash value
	hashComputer := func(filepath chan string, computedHash chan<- struct {
		HashCode
		string
	}) {
		for filePath := range filepath {
			hashCode, err := hashFile(filePath)

			if err != nil {
				hashCode = nullHash
				filePath = err.Error()
			}

			computedHash <- struct {
				HashCode
				string
			}{hashCode, filePath}
		}
	}

	// create go routines
	numCpus := runtime.NumCPU()

	if numCpus > 1 {
		numCpus = numCpus - 1
	}

	for i := 0; i < numCpus; i++ {
		go hashComputer(filepathChannel, computedHashChannel)
	}

	for _, files := range sameSizeFiles {
		hashDuplicateWithSameSize := make(map[HashCode]FileList)

		// compute hashcode and store it in map
		for _, filepath := range files {
			filepathChannel <- filepath
		}

		for i := 0; i < len(files); i++ {
			computedHash := <-computedHashChannel

			hash := computedHash.HashCode
			filepath := computedHash.string

			if hash == nullHash {
				fmt.Println("Error computing hash ", filepath)
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
		fmt.Println("Error: could not compute hashcode: ", err)
		return [20]byte{}, err
	}

	hash := sha1.Sum(data)

	return hash, nil
}
