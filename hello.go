package main

import "fmt"

const helloPrefix = "Hello, "
const holaPrefix = "Hola, "
const bonjourPrefix = "Bonjour, "

func Hello(name, language string ) string {
	if name == "" {
		name = "World"	
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
		case "Spanish":
			prefix = holaPrefix

		case "French":
			prefix = bonjourPrefix
		default:
			prefix = helloPrefix
	}
	 return
}

func main() {
	fmt.Println(Hello("world", "English"))
}