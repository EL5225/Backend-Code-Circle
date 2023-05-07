package main

import "fmt"

func luasPersegiPanjang(panjang, lebar int) int {
	luas := panjang * lebar

	return luas
}

func kelilingPersegiPanjang(panjang, lebar int) int {
	keliling := (2 * panjang) * (2 * lebar)

	return keliling
}

func volumePersegiPanjang(panjang, lebar, tinggi int) int {
	volume := panjang * lebar * tinggi

	return volume
}

func main() {
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumePersegiPanjang(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

}
