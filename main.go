package main

import "fmt"

func main() {
	var stringText string
	var intNumber int
	stringText = "Testing"
	intNumber = 1
	fmt.Println("Test World")
	fmt.Printf("%d %s\n", intNumber, stringText)
	fmt.Println(intNumber, stringText)

	floatNumber := 1.5
	fmt.Println(floatNumber)
}