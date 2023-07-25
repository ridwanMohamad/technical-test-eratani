package main

import (
	"fmt"
)

func main() {
	var testCase int
	fmt.Print("Masukan jumlah testcase (t): ")
	fmt.Scan(&testCase)
	var arrayInput []int
	fmt.Println("Input")
	for i := 0; i < testCase; i++ {
		var line int
		fmt.Scan(&line)
		arrayInput = append(arrayInput, line)
	}
	fmt.Println("Output")
	removeFirstIndex := removeFirstElement(arrayInput)
	for j := 0; j < len(removeFirstIndex); j++ {
		fmt.Println(likedInteger(removeFirstIndex[j]))
	}

}

func removeFirstElement(arr []int) []int {
	if len(arr) == 0 {
		// Array is empty, nothing to remove
		return arr
	}
	return arr[1:]
}

func isDisliked(num int) bool {
	return num%3 == 0 || num%10 == 3
}

func likedInteger(k int) int {
	count := 0
	num := 1

	for count < k {
		if !isDisliked(num) {
			count++
		}
		num++
	}

	return num - 1
}
