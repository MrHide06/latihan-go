package user

import(
	"fmt"
	"github.com/MrHide06/latihan12-interface/item"
)

type Barang struct{
	NamaBarang string
	VolumeBarang int
	HargaBarang float32
}

type User struct{
	ID int
	Username string
	Barang []Barang
}

func(u *User)Beli()(*User){
	barang := new(item.Beras)
	barang1 := new(item.Minyak)
	
	barang2 := new(item.Gula)

	// fmt.Println(user)
	fmt.Println(barang.Rice())
	fmt.Println(barang1.Oil())
	fmt.Println(barang2.Sugar())

	u.ID = 1
	u.Username = "MrHide"
	u.Barang.Barang = barang
	return u
}

func CreateUser() {
	user := new(User)

	fmt.Println("Berikut ini adalah barang yang di beli pertama kali oleh pelanggan")

	
}
