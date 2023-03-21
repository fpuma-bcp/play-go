package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(values, len(values), cap(values))
	//output: [1 2 3 4 5 6 7 8 9] 9 9

	fmt.Println(values[0])   // 1
	fmt.Println(values[:3])  // [1 2 3]
	fmt.Println(values[2:4]) // [3 4]
	fmt.Println(values[4:])  // [5 6 7 8 9]

	//Append
	values = append(values, 10, 11, 12)
	fmt.Println(values) // [1 2 3 4 5 6 7 8 9 10 11 12]

	//Add another slice
	fib_numbers := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	values = append(values, fib_numbers...)
	fmt.Println(values) //[1 2 3 4 5 6 7 8 9 10 11 12 1 1 2 3 5 8 13 21 34 55]

}
