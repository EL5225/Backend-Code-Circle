package main

import "fmt"

func main() {
	tahun := 2001
	var generasi string

	if tahun >= 1944 && tahun <= 1964 {
		generasi = "Baby Boomer"
	} else if tahun >= 1965 && tahun <= 1979 {
		generasi = "Generasi X"
	} else if tahun >= 1980 && tahun <= 1994 {
		generasi = "Generasi Y"
	} else if tahun >= 1995 && tahun <= 2015 {
		generasi = "Generasi Z"
	} else {
		generasi = "Generasi ########"
	}

	fmt.Printf("Tahun %d merupakan %s", tahun, generasi)
}
