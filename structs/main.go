package main

import "fmt"

func main() {
	var customer1 Customer
	customer1.id = 1
	customer1.name = "Ali"
	customer1.age = 30
	customer1.address = Address{city: "Istanbul", district: "Kadikoy"}
	fmt.Println(customer1)

	var customer2 = Customer{id: 2, name: "Veli", age: 40, address: Address{city: "Istanbul", district: "Besiktas"}}
	fmt.Println(customer2)

	customer3 := Customer{id: 3, name: "Deli", age: 50, address: Address{city: "Istanbul", district: "Sariyer"}}
	customer3.toString()

	changeNameByPointer(&customer3)
}

type Customer struct {
	id      int64
	name    string
	age     int
	address Address
}

type Address struct {
	city, district string
}

func (customer Customer) toString() { // Burada fonksiyonun alıcı parametresi kullanıldı. Bu sayede Customer tipi için özelleştirilmiş bir fonksiyon oluşturuldu.
	fmt.Println("Customer: ", customer.id, customer.name, customer.age)
}

// Structlar değer tiplerdir. Bu yüzden fonksiyonlara parametre olarak gönderildiğinde kopyalanır.
//Bu durumda fonksiyon içerisinde yapılan değişiklikler orjinal struct üzerinde etkili olmaz.

func changeNameByPointer(customer *Customer) {
	customer.name = "Mehmet" // Struct'ın adresi alındığı için orjinal struct üzerinde değişiklik yapılır.
}
