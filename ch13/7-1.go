package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	mutex.Lock()
	defer mutex.Unlock()
	id++
	result := id
	return result
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	max := 100
	wg.Add(max)
	for i := 0; i < max; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Finished")
}
