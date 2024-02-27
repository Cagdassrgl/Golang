package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Laptop"
	names[1] = "Mouse"
	names[2] = "Keyboard"

	fmt.Println(names) // [Laptop Mouse Keyboard]

	var prices = [3]int{1000, 50, 250}

	fmt.Println(prices) // [1000 50 250]

	var otherNames = []string{"Laptop", "Mouse", "Keyboard"} // Slices => Eleman sayısı belirtilmeyen dizi
	fmt.Println(otherNames)                                  // [Laptop Mouse Keyboard]

}
