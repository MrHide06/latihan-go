package main

import "fmt"

func main() {
	x := 4
	fmt.Println(calc(x))
}

func calc(x int) int {
	x++
	return x
}
