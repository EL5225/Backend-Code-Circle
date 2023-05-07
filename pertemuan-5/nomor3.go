package main

func buahFavorit(name string, buah ...string) string {

	var fruit string

	for _, i := range buah {
		fruit += (i + ", ")
	}

	return "halo nama saya " + name + " dan buah favorit saya adalah " + fruit
}

func main() {
	var buah = []string{"semangka", "jeruk", "melon", "pepaya", "lemon"}
	var buahFavoritJohn = buahFavorit("john", buah...)

	println(buahFavoritJohn)
}
