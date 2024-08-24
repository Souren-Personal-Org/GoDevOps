package main
import (
	"fmt"
	"sync"
)
// printWords spins off a gourinte that prints words off input channels
// until we receive a signal on exitCh. wg can be used to know that
// printWords exits.
func printWords(in1, in2 chan string, exit chan struct{}, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		select {
		case <-exit:
			fmt.Println("exiting routine")
			return
		case str := <-in1:
			fmt.Println("data from channel in1: ", str)
		case str := <-in2:
			fmt.Println("data from channel in2: ", str)
	    }
	}
}
func main() {
	in1 := make(chan string) //unbuffered channel
	in2 := make(chan string)
	wg := &sync.WaitGroup{}
	exit := make(chan struct{})
	
	wg.Add(1)
	go printWords(in1, in2, exit, wg)

	in1 <- "hello"
	in2 <- "world"
	close(exit)
	// var s struct{}
	// exit <- s
	
	wg.Wait()

}