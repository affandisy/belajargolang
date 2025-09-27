package main

import (
	"fmt"
	"sync"
	"time"
)

// Mendeklarasikan struct terlebih dahulu

type Order struct {
	Name     string
	Priority int
}

type Chef struct {
	Id   int
	Name string
}

type ResultRes struct {
	Priority int
	Message  string
}

func PriorityDispatchRestaurant(orders []Order, orderChan chan<- Order) {
	for priority := 1; priority <= 3; priority++ {
		for _, order := range orders {
			if order.Priority == priority {
				orderChan <- order
			}
		}
	}
	close(orderChan)
}

func (c Chef) assignTaskRestaurant(orderChan <-chan Order, wg *sync.WaitGroup, resultChan chan<- ResultRes) {
	defer wg.Done()

	for order := range orderChan {
		time.Sleep(3 * time.Second)

		resultChan <- ResultRes{
			Priority: order.Priority,
			Message:  fmt.Sprintf("Chef %s (%d): Cooking Order: %s", c.Name, c.Id, order.Name),
		}
	}
}

func main() {
	orders := []Order{
		{"Steak", 1},
		{"Salad", 2},
		{"Sandwich", 3},
	}

	orderChan := make(chan Order)
	resultChan := make(chan ResultRes, len(orders))

	var wg sync.WaitGroup

	chefs := []Chef{
		{1, "Gordon Ramsey"},
		{2, "Jamie Oliver"},
		{3, "Uncle Roger"},
	}

	for _, chef := range chefs {
		wg.Add(1)
		go chef.assignTaskRestaurant(orderChan, &wg, resultChan)
	}

	go PriorityDispatchRestaurant(orders, orderChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	results := make(map[int]string)

	for res := range resultChan {
		results[res.Priority] = res.Message
	}

	for i := 1; i <= len(orders); i++ {
		if msg, ok := results[i]; ok {
			fmt.Println(msg)
		}
	}

	fmt.Println("All Tasks are Completed")
}
