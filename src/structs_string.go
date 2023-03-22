package main

import "fmt"

type developer struct {
	name   string
	github string
}

func (d developer) String() string {
	return fmt.Sprintf("I am %s and my github user is %s", d.name, d.github)
}

func main() {
	myself := developer{name: "Frank", github: "fpuma-bcp"}
	fmt.Println(myself)
	// I am Frank and my github user is fpuma-bcp
}
