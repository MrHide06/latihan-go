package main

import "fmt"

type Sepeda struct{
	ban, gear int
	warna string
}

func (sepeda *Sepeda) waktu(menit float32) float32{
	
	waktuTempuh := menit * 2.5
	
	return waktuTempuh
}

func main() {
	data := [...]Sepeda{
		Sepeda{ban: 2, gear: 6, warna: "Putih"},
		Sepeda{ban: 2, gear: 5, warna: "Merah"},
		Sepeda{ban: 1, gear: 1, warna: "Hitam"},
		Sepeda{ban: 3, gear: 1, warna: "Pink"},
		Sepeda{ban: 4, gear: 2, warna: "Ungu"},
	}

	fmt.Print("Berikut ini adalah data beberapa sepeda\n")

	for _, i := range data{
		fmt.Print("==================================\n")
		fmt.Print("Warna Sepeda = ", i.warna, "\n")
		fmt.Print("Jumlah Ban Sepeda = ", i.ban, "\n")
		fmt.Print("Jumlah Gear di sepeda = ", i.gear, "\n")
		fmt.Print("20 Km waktu tempuh nya adalah ",i.waktu(20), "menit\n")
	}
}