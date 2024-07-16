package main
import "fmt"

var msg string = "hello"

func main() {
	msg := "hello from main"
	fmt.Println("man message: ", msg)
	outsider()
}

func outsider() {
	fmt.Println("package message: ", msg)
}