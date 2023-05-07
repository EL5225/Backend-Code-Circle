package main

import "fmt"

func main() {
	var angka int
	var angka2 int
	var hasil int
	var pilihan int

	fmt.Println("===== Kalkulator Sederhana ======")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Pembagian")
	fmt.Println("4. Perkalian")

	fmt.Print("Masukan pilihan : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		fmt.Print("Masukan angka 1 : ")
		fmt.Scan(&angka)
		fmt.Print("Masukan angka 2 : ")
		fmt.Scan(&angka2)

		hasil = angka + angka2
		fmt.Printf("Hasil penjumlahan : %d ", hasil)

	} else if pilihan == 2 {
		fmt.Print("Masukan angka 1 : ")
		fmt.Scan(&angka)
		fmt.Print("Masukan angka 2 : ")
		fmt.Scan(&angka2)

		hasil = angka - angka2
		fmt.Printf("Hasil pengurangan : %d ", hasil)

	} else if pilihan == 3 {
		fmt.Print("Masukan angka 1 : ")
		fmt.Scan(&angka)
		fmt.Print("Masukan angka 2 : ")
		fmt.Scan(&angka2)

		hasil = angka / angka2
		fmt.Printf("Hasil pembagian : %d ", hasil)

	} else if pilihan == 4 {
		fmt.Print("Masukan angka 1 : ")
		fmt.Scan(&angka)
		fmt.Print("Masukan angka 2 : ")
		fmt.Scan(&angka2)

		hasil = angka * angka2
		fmt.Printf("Hasil perkalian : %d ", hasil)

	} else {
		fmt.Println("Pilihan tidak ada!")
	}
}
