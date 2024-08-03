package main
import (
	"fmt"
	"strings"
)

//define a interface stringer
type Stringer interface {
	String() string
}
// type person with with its string method
type Person struct{
	First, Last string
}

func (p Person) String() string {
	return fmt.Sprintf("%s,%s", p.Last, p.First)
}

// type StrList with with its string method
type StrList []string
func (s StrList) String() string {
	return strings.Join(s, ",")
}

/*
PrintStringer() allows us to print the output of Stringer.String() of any type
that implements Stringer. Both the types we created above implement Stringer.
*/ 
func PrintStringer(s Stringer){
	fmt.Println(s.String())
}


func main() {
	var person Person = Person{First: "John", Last: "Doak"} // person := Person{First: "John", Last: "Doak"}
	var nameList Stringer = StrList{"Dave", "Sarah"} // var nameList StrList = StrList{"Dave", "Sarah"}
	PrintStringer(person)
	PrintStringer(nameList)
}