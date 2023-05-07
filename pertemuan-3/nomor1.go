package main

import "fmt"

func main() {
	var panjangPersegiPanjang int = 8
	var lebarPersegiPanjang int = 5

	var alasSegitiga int = 6
	var tinggiSegitiga int = 7

	var kelilingPersegiPanjang int
	var luasSegitiga int

	kelilingPersegiPanjang = 2 * (panjangPersegiPanjang + lebarPersegiPanjang)
	luasSegitiga = (alasSegitiga * tinggiSegitiga) / 2

	fmt.Printf("Keliling Persegi Panjang : %d \n", kelilingPersegiPanjang)
	fmt.Printf("Luas Segitiga		 : %d \n", luasSegitiga)

}
