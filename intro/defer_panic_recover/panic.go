package main
import (
	"fmt"
	// "errors"
)


func errorGen() error{
	return fmt.Errorf("error")
}

func main() {
	err := errorGen()
	fmt.Printf("%T\n", err)
	switch err.(type){
	case error:
		panic("ran into error")
	default:
		fmt.Printf("nothing")

	}
	fmt.Println("closing")
}
