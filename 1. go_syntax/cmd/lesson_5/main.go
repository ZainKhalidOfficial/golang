package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {

	// 1.
	// basicVariables()

	//2.
	// arr_slices_loops()

	//3.
	// arr_mem_allocation_speedtest()

	//4.
	strings_runes()
}

// 1.
func basicVariables() {
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

// 1.1
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

// 2.
func arr_slices_loops() {

	myIntArr := [...]int{1, 2, 3}
	fmt.Printf("Simple Int Array = %v\n", myIntArr)

	myIntSlice := []int{1, 2, 3}
	fmt.Printf("Simple Int Slice = %v\n", myIntSlice)
	fmt.Printf("Simple Int Slice Len = %v and Capacity = %v\n", len(myIntSlice), cap(myIntSlice))

	myIntSlice = append(myIntSlice, 4)
	fmt.Printf("Appended Int Slice = %v\n", myIntSlice)
	fmt.Printf("Appended Int Slice Len = %v and Capacity = %v\n", len(myIntSlice), cap(myIntSlice))

	myIntSlice2 := []int{5, 6, 7}
	fmt.Printf("Simple Int Slice2 = %v\n", myIntSlice2)
	fmt.Printf("Simple Int Slice2 Len = %v and Capacity = %v\n", len(myIntSlice2), cap(myIntSlice2))

	myIntSliceJoined := append(myIntSlice, myIntSlice2...)
	fmt.Printf("My joined int slices are = %v\n", myIntSliceJoined)
	fmt.Printf("Joined Int Slices Len = %v and Capacity = %v\n", len(myIntSliceJoined), cap(myIntSliceJoined))

	myIntSliceCopy := myIntSlice
	fmt.Printf("My Int Slice Copy = %v\n", myIntSliceCopy)

	myIntSliceCopy = append(myIntSliceCopy, 5)
	fmt.Printf("My Appended Int Slice Copy = %v\n", myIntSliceCopy)
	fmt.Printf("My Original Int Slice After It's Copy is modified = %v\n", myIntSlice)

	myIntSliceMake := make([]int, 3, 8)
	fmt.Printf("Empty Declared Slice Using Make Function = %v\n", myIntSliceMake)
	fmt.Printf("Make Slice has Len %v and Capacity %v\n", len(myIntSliceMake), cap(myIntSliceMake))

	myIntSliceMake = append(myIntSliceMake, 0)
	fmt.Printf("Make Slice after append = %v\n", myIntSliceMake)
	fmt.Printf("Make Slice after append has Len %v and Capacity %v\n", len(myIntSliceMake), cap(myIntSliceMake))

	//Map (Key-Value)
	var myMap map[string]uint = map[string]uint{"Adam": 25, "Sarah": 22, "Rehana": 23}
	fmt.Printf("\nMy Simple Map = %v\n", myMap)
	fmt.Printf("My Simple Map Indexed Output = %v\n", myMap["Sarah"])

	myMap["Alexa"] = 24
	fmt.Printf("\nMy Simple Map after direct update = %v\n", myMap)

	delete(myMap, "Sarah")

	age, exists := myMap["Sarah"]

	if exists {
		fmt.Printf("The key specified exists with age %v\n", age)
	} else {
		fmt.Println("Key does not exists")
	}

	for name, age := range myMap {
		fmt.Printf("%v has %v age\n", name, age)
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	var j = 11

	for j < 21 {
		fmt.Println(j)
		j++
	}

}

// 3.
func arr_mem_allocation_speedtest() {
	var n = 1000000

	var myIntSlice = []int{}
	var myIntSlice2 = make([]int, 0, n)

	fmt.Printf("Slice without preallocated capacity took %v to append to %v\n", timeloop(myIntSlice, n), n)

	fmt.Printf("Slice with preallocated capacity took %v to append to %v\n", timeloop(myIntSlice2, n), n)

}

// 3.1
func timeloop(mySlice []int, n int) time.Duration {

	t0 := time.Now()

	for len(mySlice) < n {
		mySlice = append(mySlice, 1)
	}

	return time.Since(t0)

}

// 4.
func strings_runes() {
	var myString = "rèsumè"
	fmt.Printf("Length of a String using len is invalid due to underlying dynamic byte encoding of utf8 = %v\n", len(myString))
	fmt.Printf("Indexing a string gives a number = %v\n", myString[0])

	for i, v := range myString {
		fmt.Printf("%v => %v\n", i, v)
	}

	// Index 2 was skipped because the è takes two bytes.

	var myStringInRuneArr = []rune(myString)
	fmt.Printf("Length of a String after converting to Rune Arr is Valid = %v\n", len(myStringInRuneArr))

	for i, v := range myStringInRuneArr {
		fmt.Printf("%v => %v\n", i, v)
	}

	//String Concatenation

	var strSlice = []string{"H", "E", "L", "L", "O"}

	var catSlice = ""
	for i := range strSlice {
		catSlice += strSlice[i]
	}

	fmt.Printf("Concatenated String is = %v\n", catSlice)
	//This is inefficient because a new arr is made on each concatenation

	//Better String Concatenation Method
	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	var catSlice2 = strBuilder.String()

	fmt.Printf("Correctly built string = %v\n", catSlice2)

}
