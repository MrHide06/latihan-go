package main

import "fmt"

func main () {
	buku := []string{"Harry-Porter", "Laskar-Pelangi", "Lord-of-The-Ring"}

	var input string
	fmt.Print("Masukan Judul BUku di sini : ")
	fmt.Scan(&input)

	key, found := Search(buku, input)
	if !found {
		fmt.Print("Buku nya belum tersedia\n")
	}else{
		fmt.Printf("Bukunya tersedia di index ke : %d\n", key)
	}
	
}

func Search(buku[]string, input string)(int, bool){
	for i, item := range buku{
		if item == input {
			return i, true
		}
	}
	return 0, false
}