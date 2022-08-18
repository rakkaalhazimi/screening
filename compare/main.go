package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/codingsince1985/checksum"
)

func contain(arr []string, content string) (string, bool) {
	var found bool
	for _, item := range arr {
		if item == content {
			found = true
			return item, found
		}
	}
	return "", found
}

func getMapdir(root string) map[string]string {
	var dirsMap = map[string]string{}
	err := filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Get path information
			dir, _ := os.Stat(path)

			// If not a directory, associate relative path with asbolute path
			if !dir.IsDir() {
				relativePath, _ := filepath.Rel(root, path)
				dirsMap[relativePath] = path
			}

			return nil
		})

	if err != nil {
		log.Println(err)
	}

	return dirsMap
}

func keyToArr(m map[string]string) []string {
	var arr []string
	for key := range m {
		arr = append(arr, key)
	}
	return arr
}

func isDifferentFile(source string, target string) bool {
	sourceHash, _ := checksum.MD5sum(source)
	targetHash, _ := checksum.MD5sum(target)

	if sourceHash == targetHash {
		return false
	} else {
		return true
	}
}

func compareDirs(source string, target string) {

	// Map of relative path : absolute path
	sourceMaps := getMapdir(source)
	targetMaps := getMapdir(target)

	// Array of relative path
	sourceDirs := keyToArr(sourceMaps)
	targetDirs := keyToArr(targetMaps)

	// Traverse all target path
	for _, sourcePath := range sourceDirs {
		foundTargetPath, isTargetFound := contain(targetDirs, sourcePath)

		if isTargetFound {
			// Compare files using absolute path
			different := isDifferentFile(sourceMaps[sourcePath], targetMaps[foundTargetPath])

			// If files differ, print modified
			if different {
				fmt.Printf("%s MODIFIED\n", sourcePath)
			}

		} else {
			// If files not found, print new
			fmt.Printf("%s NEW\n", sourcePath)
		}
	}

	// Traverse all source path
	for _, targetPath := range targetDirs {
		_, isSourceFound := contain(sourceDirs, targetPath)

		// If not found in source but target, print deleted
		if !isSourceFound {
			fmt.Printf("%s DELETED\n", targetPath)
		}
	}
}

func main() {
	var source string = "./compare/source"
	var target string = "./compare/target"
	compareDirs(source, target)

}
