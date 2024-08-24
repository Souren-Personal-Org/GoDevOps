package main
 import (
	"fmt"
 )

 func main() {
	ch := make(chan string, 1) //buffered channel with string type
	go func() {
		fmt.Println("putting data to channel")
		for _, word := range []string{"D1", "D2", "D3", "D4", "D5", "D6", "D7", "D8", "D9", "D10"} {
			ch <- word
		}
		fmt.Println("closing channel")
		close(ch)
	}()
	fmt.Println("getting data from channel")
	fmt.Println("data received(first data): ", <-ch)
	for word := range ch {
		fmt.Println("data received: ", word)
	}
 }