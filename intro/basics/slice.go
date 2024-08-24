package main
import "fmt"

func main() {
	// define slice with initial 0 value
	x := []int{} //or var x = []int
	fmt.Println("basic null slice: ", x)

	//define slice with values
	y := []int{1, 2, 3, 4}
	fmt.Println("slice size: ", len(y))

	//change index value
	y[2] = 55
	fmt.Println("new slice: ", y)
	
	//append
	y = append(y, 6, 7)
	fmt.Println("new appended slice: ", y)

	x = doAppend(x)
	fmt.Println("x changed to: ", x)

	//extract
	extractValue(y)
}


func doAppend(sl []int) []int{
	return append(sl, 100)

}

func extractValue(arr []int) {
	for index, value := range arr {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}
}