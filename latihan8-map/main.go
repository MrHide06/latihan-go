package main

import "fmt"

type produk struct{
	namaBarang string
	hargaBarang int
	stockBarang int
}
func main () {
	toko := make(map[int]produk)
	
	toko[1] = produk{namaBarang: "Oreo", hargaBarang: 10000, stockBarang: 20}
	toko[2] = produk{namaBarang: "L-Men", hargaBarang: 20000, stockBarang: 9}
	toko[3] = produk{namaBarang: "Citatos", hargaBarang: 10000, stockBarang: 10}
	toko[4] = produk{namaBarang: "Sari Roti", hargaBarang: 12000, stockBarang: 6}
	toko[5] = produk{namaBarang: "Mesis Seres", hargaBarang: 15000, stockBarang: 5}
	
	fmt.Print("Berikut ini adalah produk yang sudah mau habis :\n")
	for _, i := range toko{
		if i.stockBarang < 10 {
			fmt.Print("-",i.namaBarang, "\n")
		}
	}
}