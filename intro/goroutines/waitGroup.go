package main
import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i:=0; i<8; i++{
		wg.Add(1)
		go func (n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}
	wg.Wait()
	fmt.Println("All go routine Completed")
}