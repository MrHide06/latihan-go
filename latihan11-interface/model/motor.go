package model

import "fmt"

type Maju interface {
	Cepat()
	Lambat()
}

type Motor struct {
	Ban, Gear int
	Kecepatan float32
	Jenis     string
	Maju
}


func (m Motor) Cepat(move float32) float32 {
	fast := (move * 2.0) * 50.0
	return fast
}

func (m Motor) Lambat(move float32) float32 {
	slow := (move / 2.0) * 50.0
	return slow
}

func (m Motor) waktu(speed float32) float32{
	waktuTempuh := speed * 50.0
	
	return waktuTempuh
}

func Vehicle(){
	data := make(map[int]Motor)

	data[0] = Motor{Ban: 2, Gear: 6, Jenis: "VixionR", Kecepatan: 2.3}
	data[1] = Motor{Ban: 2, Gear: 6, Jenis: "CB150R", Kecepatan: 2.4}
	data[2] = Motor{Ban: 2, Gear: 6, Jenis: "R15", Kecepatan: 2.25}
	data[3] = Motor{Ban: 2, Gear: 6, Jenis: "CBR150R", Kecepatan: 2.5}
	data[4] = Motor{Ban: 2, Gear: 6, Jenis: "GSX150R", Kecepatan: 2.5}

	fmt.Println("Berikut ini adalah data tentang Kecepatan Motor: \n")

	for _, d := range data{
		fmt.Print("==================================\n")
		fmt.Print("Jenis Motor = ", d.Jenis, "\n")
		fmt.Print("Jumlah Ban Motor = ", d.Ban, "\n")
		fmt.Print("Jumlah Gear di sepeda = ", d.Gear, "\n")
		fmt.Print("Dalam 50 menit Kilometer yang ditempuh adalah ",d.waktu(d.Kecepatan), "Km\n")
		fmt.Print("Apabila dipercepat maka akan menembus ", d.Cepat(d.Kecepatan), "Km\n")
		fmt.Print("Apabila diperlambar maka akan menembus ", d.Lambat(d.Kecepatan), "Km\n")
	}
}