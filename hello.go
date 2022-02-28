package main

import "fmt"

const helloPrefix = "Hello, "

func Hello(name string) string {
	// string is in our domain
	return helloPrefix + name
}

func main() {
	// side effect printing to stdout
	fmt.Println(Hello("world"))
}