package main

import "fmt"

type weather struct {
	city        string
	temperature int
}

// to update the value of the parameter w
func (w *weather) incrementTemperature() {
	w.temperature++
}

// just read values, pointer is not needed
func (w weather) printValues() {
	fmt.Printf("%+v\n", w)
}

func main() {
	var myWeather weather
	myWeather.city = "Lima"
	myWeather.temperature = 23
	myWeather.printValues() // &{city:Lima temperature:23}
	myWeather.incrementTemperature()
	myWeather.printValues() // &{city:Lima temperature:24}

}
