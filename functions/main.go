package main

import "fmt"

func main() {
	var result = add(5, 6)
	fmt.Println(result)

	var result1, result2 = calculation(5, 6) // fonksiyon 2 değer döndürdüğü için 2 değişkene atandı.
	fmt.Println(result1, result2)

	var result3 = sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) // sum fonksiyonuna istediğimiz kadar değer gönderebiliriz.
	fmt.Println(result3)
}

func add(x int, y int) int {
	return x + y
}

func calculation(x int, y int) (int, int) { // fonksiyon 2 değer döndürüyor. 2'den fazla değer döndürmek istenirse parantez içinde belirtilir.
	return x + y, x - y
}

func sum(numbers ...int) int { // sum fonksiyonuna istediğimiz kadar değer gönderebiliriz. Gönderilen değerler slice olarak alınır.
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}
