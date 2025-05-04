package main

import "fmt"

func sortInts[I int8 | int16 | int32 | int64](slice []I) {
	sorted := false
	for !sorted {
		sorted = true
		for i := range slice[:len(slice)-1] {
			if slice[i] > slice[i+1] {
				sorted = false
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}
}

func main() {
	a := []int8{99, 2, 52, 100, 7, 25, 125, 35, 65}
	b := []int64{1000, 100, 25, 95, 1, 10000, 64}
	sortInts(a)
	sortInts(b)
	fmt.Println(a)
	fmt.Println(b)
}
