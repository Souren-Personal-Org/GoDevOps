package main
import "fmt"

func main() {

	//unconventional struct use
	var record = struct{
		Name string
		Age int
	}{
		Name: "Souren",
		Age: 30,
	}
	fmt.Printf("%s is %d year old\n", record.Name, record.Age)

	//type basic
	type BikeModel string
	var myBike BikeModel = "Hero" // or myBike = BikeModel("Honda")
	fmt.Println("my bike model: ", myBike)

	//conventional struct
	type Record struct{
		Name string
		Age int
	}
	var david Record = Record{Name: "David", Age: 20,}
	sarah := Record{Name: "Sarah", Age: 21,}
	fmt.Printf("%+v\n", david)
	fmt.Printf("%+v\n", sarah)

}