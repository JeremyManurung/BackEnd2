package main

import "fmt"

func main() {
	const Phi = 3.14
	var hasil,r,t float64

	fmt.Print("Masukkan Tinggi: ")
	fmt.Scanln(&t)
	fmt.Print("Masukkan jari : ")
	fmt.Scanln(&r)

	hasil = 2 * Phi * r * (r+t)

	fmt.Println("Volume balok adalah : ", hasil)

}