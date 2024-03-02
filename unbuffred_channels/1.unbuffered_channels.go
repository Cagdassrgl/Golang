package main

func main() {
	// Channel'lar go_routine'lar arasında veri alışverişi yapmamızı sağlar.

	myChannel := make(chan string) // int tipinde bir channel oluşturduk.

	go func() {
		myChannel <- "Hello World" // Channel'a veri gönderdik. // Burayı okuyan herhangi biri yoksa "Deadlock" hatası alırız.
	}()

	// Blocking => Channel'dan veri gelene kadar bekler.

	message := <-myChannel // Channel'dan veri aldık.
	println(message)
}
