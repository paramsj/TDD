package main

import "fmt"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return prefix + name
}

// no overloading, write what you need
// capital letter will export, else internal only
const prefix = "Hello, "
func HelloWithoutName() string {
	return prefix
}

func main() {
	fmt.Println(HelloWithoutName())
}