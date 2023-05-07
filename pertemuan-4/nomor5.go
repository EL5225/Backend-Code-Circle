package main

import (
	"fmt"
)

func main() {
	var satuan = map[string]int32{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
		//"Volume Balok": 168,
	}

	volumeBalok := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]

	for key, val := range satuan {
		fmt.Println(key, " = ", val)
	}

	fmt.Printf("Volume balok = %d", volumeBalok)
}
