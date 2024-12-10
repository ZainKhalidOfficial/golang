package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Printf("Hello World! I'm back...\n")

	var myInt int = 12345
	fmt.Printf("myInt is = %v\n", myInt)

	var myFloat float64 = 12345678.9
	fmt.Printf("myFloat is = %v\n", myFloat)

	var myString string = "abcdef"
	fmt.Printf("myString is =%v\n", myString)
	fmt.Printf("len(myString) is'nt always accurate due to extended range of unicode/utf8 %v\n", len(myString))
	fmt.Printf("Instead utf8.RuneCountInString is accurate = %v\n", utf8.RuneCountInString(myString))

	//Infered Var Declaration
	var myInt2 = 6789
	fmt.Printf("myInt2 without explicitly mentioning type works fine: %v\n", myInt2)

	//Shorthand Var Declaration
	myInt3 := 10111213
	fmt.Printf("myInt3 is declared in shorthand = %v\n", myInt3)

	//multiple var declaration
	myInt4, myInt5, myInt6 := 1, 2, 3
	fmt.Printf("Sum of multiple var declaration should be 6 to know that they worked = %v\n", myInt4+myInt5+myInt6)

	//const
	const pi float32 = 3.1415
	fmt.Printf("pi is an const and it's value cannot be changed after declaration = %v\n", pi)

	//function call
	result, remainder, err := intDivision(11, 10)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the division is = %v\n", result)
	// } else {
	// 	fmt.Printf("The result of the division is %v with remainder %v", result, remainder)
	// }

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the division is %v\n", result)
	default:
		fmt.Printf("The result of the division is %v with remainder %v\n", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not even close")
	}

}

func intDivision(numerator int, denumerator int) (int, int, error) {
	var err error

	if denumerator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}

	result := numerator / denumerator
	remainder := numerator % denumerator

	return result, remainder, err
}
