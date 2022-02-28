package main

import "fmt"

const helloPrefix = "Hello, "
const holaPrefix = "Hola, "

func Hello(name, language string ) string {
	if name == "" {
		name = "World"	
	}
	if language == "English" {
		return helloPrefix + name
	} else {
		return holaPrefix + name
	}
	
}

func main() {
	fmt.Println(Hello("world", "English"))
}