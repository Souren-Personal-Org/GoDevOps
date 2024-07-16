package main
import "fmt"

func main() {
	remainder, dividend := divide(12, 144)
	fmt.Printf("dividend is: %d, remainder is: %d", dividend, remainder)
}
func divide(num, div int) (res, rem int) {
	res = num / div
	rem = num % div
	return res, rem
	}