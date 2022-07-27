package main

import "fmt"
import "path/filepath"

func main() {
	var dir = "./compare/target/folder1"
	fmt.Println(filepath.Dir(dir))
}