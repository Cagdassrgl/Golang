package main

import (
	"fmt"
)

// Fonksiyon içinde fonksiyon tanımlamak
func main() {
	func() {
		fmt.Println("Hello from anonymous function")
	}()
}
