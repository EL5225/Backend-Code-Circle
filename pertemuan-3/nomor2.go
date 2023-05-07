package main

import "fmt"

func main() {

	nilai := 76

	if nilai >= 80 && nilai <= 100 {
		fmt.Println("Nilai A")
	} else if nilai >= 70 && nilai < 80 {
		fmt.Println("Nilai B")
	} else if nilai >= 60 && nilai < 70 {
		fmt.Println("Nilai C")
	} else if nilai >= 50 && nilai < 60 {
		fmt.Println("Nilai D")
	} else if nilai < 50 {
		fmt.Println("Nilai E")
	} else {
		fmt.Println("Error!")
	}
}
