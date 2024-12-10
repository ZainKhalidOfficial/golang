package main

import (
	"fmt"
)

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Println(i)
	fmt.Println(p, *p)
	*p = 10
	fmt.Println(p, *p)

	p = &i
	*p = 11

	fmt.Println(i)

	//A Usecase of pointers

	//Normally
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory address of thing1 is %p\n", &thing1)
	//To square a new mem location is used in square function
	var result = square(thing1)
	fmt.Println(result)

	//Efficient Way
	var thing3 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory address of thing3 is %p\n", &thing3)
	var result2 = efficient_square(&thing3)
	fmt.Println(result2)

}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("The memory address of temp thing2 is %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return thing2
}

func efficient_square(thing4 *[5]float64) [5]float64 {
	fmt.Printf("The memory address of temp thing4 is %p\n", thing4)
	for i := range thing4 {
		thing4[i] = thing4[i] * thing4[i]
	}

	return *thing4
}
