package main
import (
	"fmt"
	_ "sync"
)

func returnMultipleValues(n int) (string, string, int){
	return "souren", "ghosh", n
}
func main() {
	fmt.Println("this is testing")

	auto_type_assigned_var := "Souren" //auto assigned variable type
	fmt.Println(auto_type_assigned_var)

	var defined_type_assigned_var string = "Ghosh" //defined variable type
	fmt.Println(defined_type_assigned_var)

	var firstName, lastName, age := returnMultipleValues(30)
	fmt.Println(firstName, lastName, age)

}


