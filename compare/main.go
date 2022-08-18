package main

import (
	"fmt"
	"github.com/codingsince1985/checksum"
	"log"
	"os"
	"path/filepath"
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
	source_hash, _ := checksum.MD5sum(source)
	target_hash, _ := checksum.MD5sum(target)

	if source_hash == target_hash {
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

	for _, targetPath := range targetDirs {
		foundPath, isFound := contain(sourceDirs, targetPath)

		if isFound {
			// Compare files using absolute path
			different := isDifferentFile(sourceMaps[foundPath], targetMaps[targetPath])
			if different {
				fmt.Printf("%s MODIFIED", targetPath)
			}
		} else {
			fmt.Printf("%s NEW", targetPath)
		}
		fmt.Println("")
	}
}

func main() {
	var source string = "./compare/source"
	var target string = "./compare/target"
	compare_dirs(source, target)
	// fmt.Println(my_listdir1)
	// fmt.Println(my_listdir2)

}
