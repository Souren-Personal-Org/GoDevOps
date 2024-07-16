package main
import "fmt"

func main() {
	x := 90
	if x<100{
		fmt.Println("less than 90")
	}
	if res := calc_sum(); res<100{
		fmt.Println("i am ", res)
	}
	
}

func calc_sum() int{
	var x = 9
	var y = 89
	var z int
	z = x+y
	return z
}