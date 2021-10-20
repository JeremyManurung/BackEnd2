package main

import "fmt"

func main() {
	var bilanganPrima int
	tampung := 0

	fmt.Print("Masukkan Bilangan Anda : ")
	fmt.Scanln(&bilanganPrima)

	for i:=1;i<=bilanganPrima;i++{
		if bilanganPrima%i == 0{
			tampung++
		}
	}

	if tampung == 2{
		fmt.Println("Prima")
	}else{
		fmt.Println("Bukan")
	}


}