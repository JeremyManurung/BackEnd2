package main

import "fmt"

func main() {
	segitiga()
	}

func segitiga (){
	var tampung int

	fmt.Print("masukkan angka : ")
	fmt.Scanln(&tampung)

	for x :=1; x <= tampung; x++{
		for y:=tampung; y>=x; y--{
	fmt.Print(" ")
	}
	for z :=1; z<=x; z++{
	fmt.Print("* ")
}
fmt.Println()
}
}