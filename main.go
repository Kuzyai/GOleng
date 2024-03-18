package main

import "fmt"

func main() {
	var workArray [10]byte
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}
	var arr [3][2]byte
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	var change byte
	for i := 0; i < len(arr); i++ {
		change = workArray[arr[i][0]]
		workArray[arr[i][0]] = workArray[arr[i][1]]
		workArray[arr[i][1]] = change
	}
	for i := range workArray {
		fmt.Printf("%d ", workArray[i])
	}
	fmt.Printf("%T", workArray)
}
