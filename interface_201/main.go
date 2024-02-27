package main

import "fmt"

type IEncodable interface {
	Encode(value string)
	Decode(value string)
}

type XEncoder struct {
}

type YEncoder struct {
}

func (x *XEncoder) Encode(value string) {
	fmt.Println("XEncoder ile encode edildi.")
}

func (x *XEncoder) Decode(value string) {
	fmt.Println("XEncoder ile decode edildi.")
}

func (x *YEncoder) Encode(value string) {
	fmt.Println("XEncoder ile encode edildi.")
}

func (x *YEncoder) Decode(value string) {
	fmt.Println("XEncoder ile decode edildi.")
}

func main() {
	var encoder IEncodable
	encoder = &XEncoder{}
	encoder.Encode("123456")

	encoder = &YEncoder{}
	encoder.Encode("qweasd")

}
