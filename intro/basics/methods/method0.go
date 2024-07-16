package main
import "fmt"

type Record struct{
	Name string
	Age int
}
//Record's String method
func (r Record) String() string {
	return fmt.Sprintf("%s:%d", r.Name, r.Age)
}

func main() {
	var john Record = Record{Name: "John", Age: 14,} 
	fmt.Println("John's special code is-> ", john.String())
	
	//change attribute
	john.Age = 17
	fmt.Println("John's special code after change-> ", john.String())
}
