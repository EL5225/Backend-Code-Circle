package main

import "fmt"

func main() {
	tanggal := 13
	bulan := 12
	tahun := 2001
	var month string

	switch bulan {
	case 1:
		month = "Januari"
	case 2:
		month = "Februari"
	case 3:
		month = "Maret"
	case 4:
		month = "April"
	case 5:
		month = "Mei"
	case 6:
		month = "Juni"
	case 7:
		month = "July"
	case 8:
		month = "Agustus"
	case 9:
		month = "September"
	case 10:
		month = "Oktober"
	case 11:
		month = "November"
	case 12:
		month = "Desember"
	default:
		month = "########"
	}

	fmt.Printf("%d %s %d", tanggal, month, tahun)
}
