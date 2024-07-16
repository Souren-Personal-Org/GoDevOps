package main
import "fmt"

type Record struct{
	Name string
	Age int
}

//increment person age
/*
as in function and method argument are copy of reference memory, 
you need to reference pointer
*/

func (r *Record) IncrAge() {
	r.Age++ //. is a magic operator that works on struct or *struct
}

func main() {
	john := Record{Name: "John"} //age will be 0 by default
	john.IncrAge()
	fmt.Printf("%s \"s age %d", john.Name, john.Age)
}



