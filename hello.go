package main

import "fmt"

func Hello(name string) string {
	// string is in our domain
	return "Hello, " + name
}

func main() {
	// side effect printing to stdout
	fmt.Println(Hello("world"))
}