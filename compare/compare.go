package main

import "fmt"

func comparePath(arr []string, path string) bool {
	for _, value := range arr {
		if value == path {
			return true
		}
	}
	return false
}

func main() {

	var paths1 = []string{"a", "b", "c"}
	var paths2 = []string{"c", "d", "e"}

	found := comparePath(paths1, paths2[0])
	fmt.Println(found)

}
