package main

import (
	"fmt"
	"strings"
)

func main() {

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
