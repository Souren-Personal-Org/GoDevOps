package main

import "fmt"


func main() {
	args := []int{1, 2, 3, 4, 5}
	fmt.Println("sum without variadic: ", sum_without_variadic(args))
	fmt.Println("sum with variadic argument: ", sum_with_variadic(1, 2, 3, 4, 5))

}

func sum_without_variadic(numbers []int) int {
	var sum int =0
	for _, n := range numbers {
		sum += n
	}
	return sum
}


func sum_with_variadic(numbers ...int) int {
	var sum int =0
	for _, n := range numbers {
		sum += n
	}
	return sum
}