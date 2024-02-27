package main

import (
	"fmt"
)

func main() {
	pointer101()
	pointer201()
	pointer301()
}

func pointer101() {
	var a int = 10
	fmt.Println("Value of a: ", a)

	var p *int = &a // a değişkeninin adresini almak için & operatörü kullanılır.
	fmt.Println("Addres of \"a\" variable in RAM: ", p)
	fmt.Println("Value of \"a\" variable: ", *p) // a değişkeninin değerini almak için * operatörü kullanılır.

	*p = 20 // Pointer kullanarak a değişkeninin adresine ulaşıp değerini değiştirdik.
	fmt.Println("Value of \"a\" variable: ", a)
}

func pointer201() {
	var a = 10
	var b int
	var p *int

	b = a // Burada b için bellekte ayrı bir adres ayrılır ve a'nın değeri b'ye kopyalanır.
	p = &a

	*p = 20
	fmt.Println("Value of (a,b): ", a, b) // 20 10

	// Interger değer tip olduğunu için b = a dendiğinde a'nın değeri b'ye kopyalanır. Referans tip olsaydı a'nın adresi b'ye atanırdı. Bu durumda a'nın değeri değiştiğindeb'nin değeri değişirdi.
}

func pointer301() {
	var a = 10
	fmt.Println("Value of a: ", a)
	add12(10)
	fmt.Println("Value of a: ", a)

	var b = 10
	fmt.Println("Value of b: ", b)
	add12ByPointer(&b)
	fmt.Println("Value of b: ", b)

}

// pass by value
func add12(x int) {
	x = x + 12 // Burada x değişkeni kopyalandığı için a değişkeni etkilenmez.
	fmt.Println("Value of x: ", x)
}

// pass by reference
func add12ByPointer(x *int) {
	*x = *x + 12 // Burada x değişkeninin adresi alındığı için a değişkeni etkilenir.
}
