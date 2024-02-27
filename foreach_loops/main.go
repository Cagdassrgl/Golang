package main

import "fmt"

func main() {

	var numbers = []int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		fmt.Println(i, number)
	}

	var language = "Golang"

	for _, letter := range language {
		fmt.Println(letter)         // Bu şekilde yazdırılırsa ASCII değerini yazar.
		fmt.Println(string(letter)) // Bu şekilde yazdırılırsa harfi yazar.
	}
}
