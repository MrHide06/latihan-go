package main

import "fmt"

func main(){
	barang := map[int]string{
		0:"Oreo",
		1:"L-Men",
		2:"Sari Roti",
	}

	harga := map[int]int{
		0:10000,
		1:20000,
		2:12000,
	}

	stok := map[int]int{
		0:20,
		1:10,
		2:5,
	}
	fmt.Println("Berikut barang yang sudah mau habis :")
	for key, i := range stok{
		if i < 10{
			for k, b := range barang{
				for x, h := range harga{
					if (key == k) && (key == x){
						fmt.Print("-", b, " dengan harga ", h, "\n")
					}
				}
			}
		}
	}
}