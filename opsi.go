package main

import "fmt"

func main() {
	var opsi int 

	fmt.Print("masukkan opsi : ")
	fmt.Scan(&opsi)

	switch {
	case opsi == 1:
		
	const Phi = 3.14
	var hasil,r,t float64

	fmt.Print("Masukkan Tinggi: ")
	fmt.Scanln(&t)
	fmt.Print("Masukkan jari : ")
	fmt.Scanln(&r)

	hasil = 2 * Phi * r * (r+t)

	fmt.Println("Volume balok adalah : ", hasil)


	case opsi == 2:

	default:
		fmt.Print("invalid option")

	}
	
}
