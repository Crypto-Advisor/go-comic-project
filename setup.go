package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("What is your name?:")
	fmt.Scan(&name)
	fmt.Println("Hello", name)

	var age int
	fmt.Println("How old are you?:")
	fmt.Scan(&age)
	fmt.Printf("Hello %s you are %d years old\n", name, age)

	var newMessage = fmt.Sprintf("whats up %s you look like your %d years old", name, age)
	fmt.Println(newMessage)

}
