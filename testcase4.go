package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Generate Int
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(10)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			//Swap if j more than j+1
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println(arr)
	}

	fmt.Println("Final Sort : ", arr)
}
