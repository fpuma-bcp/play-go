package main

import (
	"example/playground/basic/src/entity"
	"fmt"
)

func main() {
	var myPhone entity.Device
	myPhone.TypeDevice = "SmartPhone"
	myPhone.Brand = "Nokia"
	fmt.Println(myPhone) // {SmartPhone Nokia  0}
}
