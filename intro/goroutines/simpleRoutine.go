package main

import (
	"fmt"
)

func main() {
	for i := 0; i<10; i++{
		x := i
		go fmt.Printf("%d\n", x)
	}
	fmt.Println("go routine are on go")
	select {}
}