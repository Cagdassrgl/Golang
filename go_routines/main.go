package main

import (
	"fmt"
	"sync"
)

// main bir go_routine'dir yani thread
func main() {
	// go f1() Ayrı bir thread'de çalıştırır
	// go f2()
	// time.Sleep(3 * time.Second) 3 saniye bekler

	wg := sync.WaitGroup{}
	wg.Add(2) // 2 adet fonksiyonu bekleyeceği için 2 verdik

	go func() {
		defer wg.Done() // f1 fonksiyonu bittiğinde 1 azaltır
		f1()
	}()
	go func() {
		defer wg.Done() // f2 fonksiyonu bittiğinde 1 azaltır
		f2()
	}()

	wg.Wait() // 0 olana kadar bekler

	fmt.Print("End of the main function")

}

func f1() {
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}

// Kendi go_routine'umuzu oluşturabiliriz. Bu sayede main go_routine'umuzun işi bitse bile bu go_routine çalışmaya devam eder.
