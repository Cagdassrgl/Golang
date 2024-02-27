package main

import "fmt"

func main() {
	names := make(map[string]int)

	names["John"] = 21
	names["Steve"] = 30
	names["Sam"] = 29

	var ages = map[string]int{
		"John":  21,
		"Steve": 30,
		"Sam":   29,
	}

	fmt.Println(ages) // map[John:21 Steve:30 Sam:29]

	// fmt.Println(names) // map[John:21 Steve:30 Sam:29]

	fmt.Println(names["John"])   // 21
	fmt.Println(names["Cagdas"]) // 0 Eğer olmayan bir key'i çağırırsak 0 döner.

	delete(names, "Sam") // Sam key'ini siler.
}
