package main

import "fmt"

func changeValueAtZeroIndex(arr [2]int){
	arr[0] = 3
	fmt.Println("Inside function: ", arr[0])
}

func main() {
	x := [2]int{}
	changeValueAtZeroIndex(x)
	fmt.Println(x)
}