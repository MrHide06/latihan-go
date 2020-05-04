package sepeda

import "fmt"

type Travelling interface{
	waktu()
}

type Sepeda struct{
	ban, gear int
	warna string
	kecepatan float32
}

func (sepeda *Sepeda) waktu(speed float32) float32{
	
	waktuTempuh := speed * 2.5
	
	return waktuTempuh
}

func (s Sepeda) Cepat(move float32) float32 {
	fast := (move * 2.0) * 50.0
	return fast
}

func (s Sepeda) Lambat(move float32) float32 {
	slow := (move / 2.0) * 50.0
	return slow
}

func Bike() {
	data := make(map[int]Sepeda)

	data[0] = Sepeda{ban: 2, gear: 6, warna: "Putih", kecepatan: 0.4}
	data[1] = Sepeda{ban: 2, gear: 5, warna: "Merah", kecepatan: 0.3}
	data[2] = Sepeda{ban: 1, gear: 1, warna: "Hitam", kecepatan: 0.1}
	data[3] = Sepeda{ban: 3, gear: 1, warna: "Pink", kecepatan: 0.1}
	data[4] = Sepeda{ban: 4, gear: 2, warna: "Ungu", kecepatan: 0.2}

	fmt.Print("Berikut ini adalah data beberapa sepeda\n")

	for _, i := range data{
		fmt.Print("==================================\n")
		fmt.Print("Warna Sepeda = ", i.warna, "\n")
		fmt.Print("Jumlah Ban Sepeda = ", i.ban, "\n")
		fmt.Print("Jumlah Gear di sepeda = ", i.gear, "\n")
		fmt.Print("Dalam 50 menit Kilometer yang ditempuh adalah ",i.waktu(i.kecepatan), "Km\n")
		fmt.Print("Apabila dipercepat maka akan menembus ", i.Cepat(i.kecepatan), "Km\n")
		fmt.Print("Apabila diperlambar maka akan menembus ", i.Lambat(i.kecepatan), "Km\n")
	}
}