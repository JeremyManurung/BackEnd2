package main

import "fmt"

func main() {
   // input
	var nilai int

	fmt.Print("Masukkan Nilai Anda : ")
	fmt.Scanln(&nilai)

	if nilai >= 80 {
		fmt.Print("A")
	}else if nilai >= 69{
		fmt.Print("B")
	}else if nilai >= 50{
		fmt.Print("C")
	}else if nilai >= 35{
		fmt.Print("D")
	}else if nilai >= 0{
		fmt.Print("E")
	}else{
		fmt.Print("Invalid")
	}
}