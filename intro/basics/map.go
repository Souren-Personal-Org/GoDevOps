package main

import "fmt"

func main() {
	//first way to create map
	var counterMap = make(map[string]int, 10)
	counterMap["pages"] = 10
	fmt.Println("counter map is: ", counterMap)

	//conventional way
	var itemToType = map[string]string{
		"fruit": "banana","drinks": "coffee",
	}
	var fruitName string = itemToType["fruit"]
	fmt.Println("fruit name is: ", fruitName)

	// test if a item exists
	itemToType["vegetable"] = "cabbage"
	if vegetable, ok := itemToType["vegetable"]; ok{
		fmt.Printf("vegetable is found %q", vegetable)
	}else {
		fmt.Println("vegetable is not found")
	}

	//extracting items
	fmt.Println("")
	for k, v := range itemToType {
		fmt.Printf("extracted item key: %q value: %q\n", k, v)
	}

}