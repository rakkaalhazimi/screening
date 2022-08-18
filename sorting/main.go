package main

import "fmt"

// import "strings"

func max(arr []int) int {
	var max_value int = 0
	for _, value := range arr {
		if max_value < value {
			max_value = value
		}
	}
	return max_value
}

func vertical_bar(arr []int) {
	var max_value int = max(arr)
	for i := max_value; i > 0; i-- {
		for j := 0; j < len(arr); j++ {
			if arr[j] >= i {
				fmt.Print(" |")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
	for _, value := range arr {
		fmt.Printf(" %d", value)
	}
	fmt.Println("")
}

func insertion_sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				vertical_bar(arr)
				fmt.Println("")
			}
		}
	}
	return arr
}

func reverse_insertion_sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				vertical_bar(arr)
				fmt.Println("")
			}
		}
	}
	return arr
}

func main() {
	var arr = []int{2, 1, 3, 4, 6, 5, 9, 7}
	// Gunakan insertion_sort atau reverse_insertion_sort
	var result = insertion_sort(arr)
	fmt.Println(result)
}
