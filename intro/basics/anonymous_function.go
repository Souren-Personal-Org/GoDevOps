package main
import "fmt"

func main() {
	var result string = func (word1 string, word2 string) string {
		return word1 + " " + word2
	}("anonymous", "function")
	fmt.Println(result)
}