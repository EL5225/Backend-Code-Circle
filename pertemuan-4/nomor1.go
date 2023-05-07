package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 30; i++ {
		if i%2 == 1 {
			if i%3 == 0 {
				fmt.Printf("%d - I Love Coding\n", i)
			} else {
				fmt.Printf("%d - CODE\n", i)
			}
		} else if i%2 == 0 {
			fmt.Printf("%d - CIRCLE\n", i)
		}
	}

}
