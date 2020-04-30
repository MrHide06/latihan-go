package main

import (
	"errors"
	"fmt"
)

// type Names struct {
// 	firstName string
// 	lastName  string
// }

type Person struct {
	Name string
	age  int
}

func (e *Person) sayHello(halo string) string {
	return e.Name + halo
}

func (e *Person) Berjalan() string {
	return "berjalan"
}

func (e *Person) Berbicara() string {
	return "Halloooo"
}

func (e *Person) ChangeName(name string) (bool, error) {
	if name == "" {
		return false, errors.New("Error tidak bisa ganti karena kosong")
	}
	e.Name = name
	return true, nil
}

type Anak struct{
	Person
	School string
}

func (a * Anak) SetSchool(name string) error{
	if name == ""{
		return errors.New("Errors")
	}
	a.School = name
	return nil
}

func main() {
	// p := Person{Name: "rubi", age: 20}

	// ok, err := p.ChangeName("")
	// if ok {
	// 	fmt.Println("Pergantian nama berhasil")
	// } else {
	// 	fmt.Println(err)
	// }
	// fmt.Println(p)

	// p.Names = Names{firstName: "Rubi", lastName: "Anggoro"}

	// fmt.Println(p)
	// fmt.Println(p.age)
	// // fmt.Println(p.Names)

	// fmt.Println(p.sayHello(" selamat pagi"))
	// fmt.Println(p.Berjalan())
	// fmt.Println(p.Berbicara())

	a := Anak{}
	fmt.Println(a)
	err := a.SetSchool("SMKN 10 JAKARTA")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

	ok, err := a.ChangeName("Wahid")
	if !ok {
		fmt.Println(err)
	}
	fmt.Println(a)
	
}