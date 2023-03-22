package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Rectangle struct {
	base   float64
	height float64
}

func (rectangle Rectangle) getArea() float64 {
	return rectangle.base * rectangle.height
}

type Square struct {
	side float64
}

func (square Square) getArea() float64 {
	return square.side * square.side
}

func printArea(shape Shape) {
	fmt.Printf("%+v\n", shape)
	fmt.Println("area:", shape.getArea())

}
func main() {
	mySquare := Square{side: 4}
	myRectangle := Rectangle{base: 2, height: 4}
	printArea(mySquare)
	printArea(myRectangle)

}
