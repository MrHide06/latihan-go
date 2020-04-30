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
	data := make(map[int]Sepeda)

	data[0] = Sepeda{ban: 2, gear: 6, warna: "Putih"}
	data[1] = Sepeda{ban: 2, gear: 5, warna: "Merah"}
	data[2] = Sepeda{ban: 1, gear: 1, warna: "Hitam"}
	data[3] = Sepeda{ban: 3, gear: 1, warna: "Pink"}
	data[4] = Sepeda{ban: 4, gear: 2, warna: "Ungu"}

	fmt.Print("Berikut ini adalah data beberapa sepeda\n")

	for _, i := range data{
		fmt.Print("==================================\n")
		fmt.Print("Warna Sepeda = ", i.warna, "\n")
		fmt.Print("Jumlah Ban Sepeda = ", i.ban, "\n")
		fmt.Print("Jumlah Gear di sepeda = ", i.gear, "\n")
		fmt.Print("Sepeda ini apabila menempuh jarak 20 km akan menghabiskan waktu sekitar ",i.waktu(20), "menit\n")
	}
}