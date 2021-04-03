package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	a := 0.0
	var wg sync.WaitGroup

	var mu sync.Mutex // guards access

	wg.Add(10000000)
	for i := 0; i < 10000000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			a += math.Sqrt(float64(i * 3))
		}()
	}
	wg.Wait()
	fmt.Println(a) // will always be about 3.6515~e+10
}
