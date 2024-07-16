package main
import "fmt"


func main() {
	// memory address
	var x int = 4
	fmt.Println("memory addr of x var: ", &x)

	//declare pointer

	var intPtr *int
	intPtr = &x
	fmt.Println("pointer value is same as x's memory addr: ", intPtr)

	//dereferencing
	fmt.Println("value stored in intPtr: ", *intPtr)

	//change pointer value
	*intPtr = 12
	fmt.Println("changed value of x: ", x)

	//change in function
	changeValue(intPtr)
	fmt.Println("function has changed me to: ", x)
}

func changeValue(p *int) {
	*p += 10

}