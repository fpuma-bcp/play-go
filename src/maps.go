package main

import "fmt"

func main() {
	currency := make(map[string]string)
	currency["en"] = "dollar"
	currency["pe"] = "sol"
	currency["br"] = "real"

	fmt.Println(currency) // map[br:real en:dollar pe:sol]

	for k, v := range currency {
		fmt.Println(k, "--", v)
	}

	//Find value
	chile_currency, cl_found := currency["cl"]
	fmt.Println(chile_currency, cl_found)
	brazil_currency, br_found := currency["br"]
	fmt.Println(brazil_currency, br_found)
}
