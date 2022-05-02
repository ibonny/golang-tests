package main

import "fmt"

type mystuff struct {
	value string
}

func (ms *mystuff) testfunc() {
	fmt.Println(ms.value)
}

func main() {
	fmt.Println("This is a test.")

	ms := mystuff{"some string"}

	ms.testfunc()
}
