package main

import (
	"fmt"
)

type Makanan struct {
	Name string
}

type Mamalia interface {
	Move() string
	Eat() Makanan
}

type Person struct{}

func (u *Person) Move() string {
	return "Walk"
}

type IkanPaus struct{}

func (i *IkanPaus) Move() string {
	return "Swim"
}

func CheckCanSwim(m Mamalia) (string, bool) {
	if m.Move() == "Walk" {
		return m.Move(), false
	}

	return m.Move(), true
}

func main() {
	fmt.Println(CheckCanSwim(&Person{}))
	fmt.Println(CheckCanSwim(&IkanPaus{}))
}