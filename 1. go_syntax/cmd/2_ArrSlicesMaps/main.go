package main

import (
	"fmt"
	"time"
)

func main() {

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

	arr_mem_allocation_speedtest()

	//A copied slice also keep reference to the original slice
	var mySliceToCopy = []int{1, 2, 3}
	var mySliceCopied = mySliceToCopy
	mySliceCopied[2] = 4
	fmt.Println(mySliceToCopy)

}

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
