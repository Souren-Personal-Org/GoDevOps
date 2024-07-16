package main
import "fmt"

func main() {
	var i int

	for i<10 {
		i++
	}
	x := true

	for x { //loop forever
		fmt.Println("forever loop")
	}
}