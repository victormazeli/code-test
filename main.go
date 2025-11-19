package main

import (
	"fmt"
	"sync"
	"time"
)

var counter string
var wg = sync.WaitGroup{}
var numberOfWorkers = 3
var mu = sync.Mutex{}

func runWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting worker !\n")
	time.Sleep(time.Second)
	fmt.Printf("Printing id output: %d, \n", id)

}

// func compute(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Starting compute !\n")
// 	counter = fmt.Sprintf("compute id of %d", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Printing counter output: %v, \n", counter)
// }

// func computeWithMutex(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Starting compute !\n")
// 	mu.Lock()
// 	counter = fmt.Sprintf("compute id of %d", id)
// 	mu.Unlock()
// 	time.Sleep(time.Second)
// 	fmt.Printf("Printing counter output: %v, \n", counter)
// }

func main() {
	for i := 1; i <= numberOfWorkers; i++ {
		wg.Add(1)
		go runWorker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All jobs done!")

}

// func main() {

// 	wg.Add(numberOfWorkers)

// 	for i := 1; i <= numberOfWorkers; i++ {
// 		go runWorker(i, &wg)
// 	}

// 	wg.Wait()

// 	fmt.Println("All jobs done!")

// }

// func main() {
// 	wg.Add(numberOfWorkers)
// 	for i := 0; i <= numberOfWorkers; i++ {
// 		go runWorker(i, &wg)
// 	}

// 	wg.Wait()

// 	fmt.Println("All jobs done!")

// }

// func main() {
// 	for i := 1; i <= numberOfWorkers; i++ {
// 		wg.Add(1)
// 		go compute(i, &wg)
// 	}

// 	wg.Wait()

// 	fmt.Println("All jobs done!")

// }

// func main() {
// 	for i := 1; i <= numberOfWorkers; i++ {
// 		wg.Add(1)
// 		go computeWithMutex(i, &wg)
// 	}

// 	wg.Wait()

// 	fmt.Println("All jobs done!")

// }
