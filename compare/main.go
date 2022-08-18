package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func contain(arr []string, content string) bool {
	var found bool = false
	for i := 0; i < len(arr); i++ {
		if arr[i] == content {
			found = true
			break
		}
	}
	return found
}

func get_listdir(root string) []string {
	var dirs = []string{}
	err := filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			dir, _ := os.Stat(path)

			// If not a directory, get the relative path
			if !dir.IsDir() {
				relative_path, _ := filepath.Rel(root, path)
				dirs = append(dirs, relative_path)
			}

			return nil
		})

	if err != nil {
		log.Println(err)
	}
	return dirs
}

func compare_dirs(source string, target string) {
	source_dirs := get_listdir(source)
	target_dirs := get_listdir(target)

	for i := 0; i < len(target_dirs); i++ {
		var found bool = contain(source_dirs, target_dirs[i])

		if found {
		} else {
			fmt.Printf("%s NEW", target_dirs[i])
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
