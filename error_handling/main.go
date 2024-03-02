package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	file, error := os.ReadFile("sample.txt")

	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("File content: ", string(file))
	}

	result, err := divide(100, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}

}

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	} else {
		return x / y, nil
	}
}
