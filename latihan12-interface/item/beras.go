package item

import "fmt"

type Beras struct{
	Nama string
	Volume int
	Harga float32	
}

func (b *Beras)Rice()(*Beras){
	b.Nama = "Beras"
	b.Volume = 10
	b.Harga = 120000.0

	fmt.Println(b.Nama)

	return b
}
