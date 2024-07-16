package main
import "fmt"

func main() {
	var x int = 10
	response := newWorld(2000, 24)
	fmt.Println("newWorld returned: ", response)
	switch year := newWorld(2000, 24);{
	case year>2024 && year<2035:
		fmt.Println("Welcome to new year")
	case (year == 2024):
		fmt.Println("Welcome to year 2024")
	default:
		fmt.Println("Year is out of range")
	}

	switch x {
	case 3:
		fmt.Println("x is 3")
	case 10, 15:
		fmt.Println("x is either 10 or 15")
	default:
		fmt.Println("x is unknown")

	}
}

func newWorld(a int, b int) int{
	c := a+b
	return c

}