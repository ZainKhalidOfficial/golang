package main

import (
	"fmt"
)

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
	brand
}

type owner struct {
	name string
}

type brand struct {
	model string
}

func (e gasEngine) milesleft() uint8 {
	return e.gallons * e.mpg
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) milesleft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesleft() uint8
}

func main() {

	var myEngine gasEngine = gasEngine{mpg: 50, gallons: 15, ownerInfo: owner{"Alex"}, brand: brand{"XYZ"}}
	myEngine.mpg = 40
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name, myEngine.model)
	fmt.Printf("We can travel %v miles \n", myEngine.milesleft())
	canWeMakeIt(myEngine, 88)

	var myEngine2 electricEngine = electricEngine{40, 15}
	canWeMakeIt2(myEngine2, 88)

	//Anonymous Structs (Not Reusable)

	var myEngine3 = struct {
		mpg     uint8
		gallons uint8
	}{32, 10}

	fmt.Println(myEngine3.mpg, myEngine3.gallons)

}

func canWeMakeIt(e gasEngine, miles uint8) {
	if e.milesleft() >= miles {
		fmt.Println("Yes We Can Make It")
	} else {
		fmt.Println("No, We can't make it")
	}
}

func canWeMakeIt2(e engine, miles uint8) {
	if e.milesleft() >= miles {
		fmt.Println("Yes We Can Make It")
	} else {
		fmt.Println("No, We can't make it")
	}
}
