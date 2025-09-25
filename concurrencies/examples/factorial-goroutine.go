package main

import (
	"fmt"
	"sync"
)

func Factorial(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	result := 1
	for i := 2; i <= n; i++ {
		result = result * i
	}

	fmt.Printf("Factorial of %d is %d \n", n, result)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var wg sync.WaitGroup

	wg.Add(len(numbers))

	for _, num := range numbers {
		go Factorial(num, &wg)
	}

	wg.Wait()
	fmt.Println("All Factorial is Calculated")
}
