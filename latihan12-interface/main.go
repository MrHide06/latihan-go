package main

import(
	"fmt"
	"github.com/MrHide06/latihan12-interface/item"
)

type User struct{
	ID int
	Username, Fullname, Email, Password string
}

func(u *User) Beli()(int, string){
	buy := new(item)
}