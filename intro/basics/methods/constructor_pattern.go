package main
import "fmt"

type Record struct{
	Name string
	Age int
}

// validate record constructor pattern
func NewRecord(name string, age int) (*Record, error) {
	if name == "" {
		return nil, fmt.Errorf("name can not be empty string")
	}
	if age <= 0 {
		return nil, fmt.Errorf("age can not be <0")
	}
	return &Record{Name: name, Age: age}, nil
}


func main() {
	rec, err := NewRecord("John", 100)
	fmt.Println(rec)
	if err != nil {
		fmt.Errorf("error occurred")
	}
}

