package main

import "fmt"

func introduce(name, gender, job, age string) (perkenalan string) {

	if gender == "laki-laki" {
		perkenalan = "Pak " + name + " adalah seorang " + job + " berusia " + age + " tahun"
	} else if gender == "perempuan" {
		perkenalan = "Bu " + name + " adalah seorang " + job + " berusia " + age + " tahun"
	}

	return
}

func main() {
	john := introduce("john", "laki-laki", "penulis", "30")
	sarah := introduce("Sarah", "perempuan", "model", "28")

	fmt.Println(john)
	fmt.Println(sarah)

}
