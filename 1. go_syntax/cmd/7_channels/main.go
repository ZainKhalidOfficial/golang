package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 5

func main() {
	//1. Basic Example
	// var c = make(chan int, 5)
	// go process(c)

	// for i := range c {
	// 	fmt.Println(i)
	// 	time.Sleep(1 * time.Second)
	// }

	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	var chicken_deal_channel = make(chan string)
	var tofu_deal_channel = make(chan string)

	for i := range websites {
		go checkWebsiteForChicken(websites[i], chicken_deal_channel)
		go checkWebsiteForTofu(websites[i], tofu_deal_channel)
	}

	sendMessage(chicken_deal_channel, tofu_deal_channel)

}

//1.
// func process(c chan int) {
// 	defer close(c)
// 	for i := 0; i < 5; i++ {
// 		c <- i
// 	}
// 	fmt.Println("Existing Function")
// }

func checkWebsiteForChicken(website string, good_deal chan string) {

	for {
		time.Sleep(time.Second * 1)
		var price = rand.Float32() * 20
		if price < MAX_CHICKEN_PRICE {
			good_deal <- website
			break
		}
	}
}

func checkWebsiteForTofu(website string, good_deal chan string) {

	for {
		time.Sleep(time.Second * 1)
		var price = rand.Float32() * 20
		if price < MAX_TOFU_PRICE {
			good_deal <- website
			break
		}
	}
}

func sendMessage(chicken_deal_channel chan string, tofu_deal_channel chan string) {
	// fmt.Printf("Find a good deal at %s\n", <-chicken_deal_channel)

	select {
	case website := <-chicken_deal_channel:
		fmt.Printf("Found a Deal On Chicken at %v", website)
	case website := <-tofu_deal_channel:
		fmt.Printf("Found a Deal On Tofu at %v", website)
	}
}
