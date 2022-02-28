package main

import "fmt"

func Hello() string {
	// string is in our domain
	return "Hello, world"
}

func main() {
	// side effect printing to stdout
	fmt.Println(Hello())
}