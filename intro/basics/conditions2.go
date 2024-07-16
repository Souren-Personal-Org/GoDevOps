package main
import "fmt"

func main() {
	param1 := 10
	param2 := 12
	if res := sum_up(param1, param2); res<34 {
		fmt.Println("less than 34:: ", res)
	} else if r := sum_up(param1, param2); r<34 {
		fmt.Println("in between 30-34", r)

	} else {
		fmt.Println("Outside of range")
	}
}

func sum_up(a int, b int) int {
	return a+b
}