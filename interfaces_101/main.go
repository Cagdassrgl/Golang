package main

type IShippable interface {
	shippingCost() int
}

type Book struct {
	desi int
}

type Electronic struct {
	desi int
}

type Flower struct {
	desi int
}

func (b *Book) shippingCost() int {
	return 5 + (b.desi * 2)
}

func (e *Electronic) shippingCost() int {
	return 10 + (e.desi * 3)
}

func (f *Flower) shippingCost() int {
	return 10 + (f.desi * 3)
}

func calculateTotalShippingCost(products ...IShippable) int {
	var total int
	for _, product := range products {
		total += product.shippingCost()
	}
	return total
}

func main() {

	var product, product2, product3 IShippable
	product = &Book{desi: 10}
	println(product.shippingCost())
	product2 = &Electronic{desi: 20}
	println(product2.shippingCost())
	product3 = &Flower{desi: 30}
	println(product3.shippingCost())

	totalShippingCost := calculateTotalShippingCost(product, product2)
	println(totalShippingCost)

}
