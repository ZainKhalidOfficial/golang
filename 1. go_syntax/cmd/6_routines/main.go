package main

import (
	"fmt"
	"sync"
	"time"
)

// var m = sync.Mutex{}
var rm = sync.RWMutex{}

var wg = sync.WaitGroup{}
var database = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {

	var v0 = time.Now()

	for i := range database {
		wg.Add(1)
		go dbCall(i)
	}

	//performance test
	// for i := 0; i < 1000; i++ {
	// 	go count()
	// }

	wg.Wait()

	fmt.Printf("\nTotal time to execution = %v\n", time.Since(v0))
	fmt.Printf("\nDB contains following data = %v\n", results)
}

func dbCall(i int) {

	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The result of dbCall %v is %v\n", i+1, database[i])
	// m.Lock()
	// results = append(results, database[i])
	// m.Unlock()

	save(database[i])
	log()
	wg.Done()
}

func save(result string) {
	rm.Lock()
	results = append(results, result)
	rm.Unlock()
}

func log() {
	rm.RLock()
	fmt.Printf("The current results are = %v\n", results)
	rm.RUnlock()
}

// performance test
func count() {
	var i int = 0
	for i < 1000000000 {
		i++
	}
}
