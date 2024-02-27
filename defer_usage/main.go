package main

import "fmt"

func main() {
	// Defer anahtar kelimesi bir işlemi ertelemeye yarar. await'e benzer.
	// defer fmt.Println("1. işlem")
	// defer fmt.Println("2. işlem")
	// defer fmt.Println("3. işlem")

	// fmt.Println("Main")

	example()

	var condition bool = true

	defer cleanup()

	if condition {
		panic("Bir hata oluştu.") // Bu anahtar kelime ile programın çalışması tamamen durdurulur. Altındaki hiçbir şey çalışmaz.
	}
	// cleanup() // Bu kod çalışmaz. Çalışması için panic'ten önce defer ile çağrılmalıdır.

	// Çıktı: Main 3. işlem 2. işlem 1. işlem

	// Defer'ler Stack mantığında çalışır. Yani en son eklenen en önce çalışır. FILO (First In Last Out)
}

func example() {
	var condition bool = true
	defer fmt.Println("1. işlem")
	if condition {
		return
	}
	defer fmt.Println("2. işlem")

	// Çıktı: 1. işlem
}

func cleanup() {
	fmt.Println("###### Cleanup worked ######")
}
