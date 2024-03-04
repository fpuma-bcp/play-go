package main

import "fmt"

func main() {
	fmt.Println("Hello")
	go func(text string) {
		fmt.Println(text)
	}("World")

	defer fmt.Println("Bye")
}

//Output
//World
//Hello
