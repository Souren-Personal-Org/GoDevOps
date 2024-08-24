package main
import "fmt"

func printStuff() (v string) {
	defer fmt.Println("exiting")
	defer func() {
		v = "releasing lock in the file"
	}()
	fmt.Println("opening and writing to a file")
	return 
}
func main() {
	var v string = printStuff()
	fmt.Println(v)
}