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

// Dispatcher mengirim order ke channel sesuai urusan prioritas
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

// Chef mengambil order dari channel, lalu mengirim hasil ke resultChan
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
	// Data Pesanan
	orders := []Order{
		{"Steak", 1},
		{"Soup", 1},
		{"Burger", 2},
		{"Salad", 2},
		{"Sandwich", 3},
	}

	// Channel komunikasi
	orderChan := make(chan Order)
	resultChan := make(chan ResultRes, len(orders))

	var wg sync.WaitGroup

	// Data Chef
	chefs := []Chef{
		{1, "Gordon Ramsey"},
		{2, "Jamie Oliver"},
		{3, "Uncle Roger"},
	}

	// Jalankan Chef
	for _, chef := range chefs {
		wg.Add(1)
		go chef.assignTaskRestaurant(orderChan, &wg, resultChan)
	}

	// Jalankan Dispatcher
	go PriorityDispatchRestaurant(orders, orderChan)

	// Tutup resultChan setelah semua chef selesai
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// results := make(map[int]string)

	// Gunakan map[int][]string agar bisa menampung lebih dari satu hasil tiap prioritas
	results := make(map[int][]string)

	for res := range resultChan {
		// results[res.Priority] = res.Message
		results[res.Priority] = append(results[res.Priority], res.Message)
	}

	// Cetak hasil sesuai urutan prioritas
	for i := 1; i <= len(orders); i++ {
		if msgs, ok := results[i]; ok {
			for _, msg := range msgs {
				fmt.Println(msg)
			}
		}
	}

	fmt.Println("All Tasks are Completed")
}
