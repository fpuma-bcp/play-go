package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar) // {Ford 2020}

	var luxuryCar car
	luxuryCar.brand = "Ferrari"
	fmt.Println(luxuryCar) // {Ferrari 0}

}
