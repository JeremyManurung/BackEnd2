package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan Bilangan : ")
	fmt.Scanf("%d", &bilangan)
	tablePerkalian(bilangan)
}

func tablePerkalian (bilangan int){
	for i := 1; i<=bilangan;i++{
		for j:=1;j<=bilangan;j++{
			fmt.Printf("%d", j*i);
		}
		fmt.Println("")
	}
}