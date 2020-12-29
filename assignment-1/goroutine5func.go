package main

import (
	"fmt"
)

func main() {
	go printSomething()
	fmt.Print("1")
}

func printSomething() {
	fmt.Print("0")
}