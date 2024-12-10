package main

import (
	"fmt"
)

func main() {
	var p *int32 = new(int32)
	var i int32

	fmt.Println(p, *p, &i, i)
	*p = 10
	fmt.Println(p, *p, &i, i)
	p = &i
	fmt.Println(p, *p, &i, i)
	i = 11
	fmt.Println(p, *p, &i, i)

	//UseCase

	//Normal Squaring (Uses another mem address during squaring func)
	var thing1 = [5]float32{1, 2, 3, 4, 5}
	fmt.Println("\n\nthing1 = ", thing1)
	fmt.Printf("Address of thing 1 = %p\n", &thing1)

	var result1 = square(thing1)
	fmt.Println("result1 = ", result1)

	//Efficient Squaring using Pointers
	var thing3 = [5]float32{1, 2, 3, 4, 5}
	fmt.Println("\n\nthing3 = ", thing3)
	fmt.Printf("Address of thing 3 = %p\n", &thing3)

	var result2 = efficient_square(&thing3)
	fmt.Println("result2 = ", result2)
}

func square(thing2 [5]float32) [5]float32 {
	fmt.Printf("Address of thing 2 = %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

func efficient_square(thing4 *[5]float32) [5]float32 {
	fmt.Printf("Address of thing 4 = %p\n", thing4)
	for i := range thing4 {
		thing4[i] = thing4[i] * thing4[i]
	}
	return *thing4
}
