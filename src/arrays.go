package main

import "fmt"

func main() {
	var numbers [4]int
	numbers[0] = 100
	numbers[1] = 101
	numbers[2] = 102
	fmt.Println(numbers, len(numbers), cap(numbers))
	//output: [100 101 102 0] 4 4

}
