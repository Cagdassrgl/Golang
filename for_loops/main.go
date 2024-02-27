package main

import "fmt"

func main() {

	index := 1

	for index <= 10 {
		fmt.Println(index)
		index++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 1; i <= 10; i++ {
		if i == 3 {
			break // i değeri 3 olduğunda döngüyü sonlandırır.
		}
		fmt.Println(i) // 1, 2
	}

	for i := 1; i <= 10; i++ {
		if i == 3 {
			continue // i değeri 3 olduğunda devamındaki kodları çalıştırmaz.
		}
		fmt.Println(i) // 1, 2, 4, 5, 6, 7, 8, 9, 10
	}
}
