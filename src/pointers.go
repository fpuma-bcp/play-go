package main

import "fmt"

func main() {
	a := 500
	b := &a

	fmt.Println(a) // 500

	fmt.Println(b)  // 0xc000114008
	fmt.Println(*b) // 500

	*b = 200
	fmt.Println(a)  // 200
	fmt.Println(*b) // 200

}
