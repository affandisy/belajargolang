package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type CustomerReal struct {
	Id   int
	Name string
}

type OrderReal struct {
	Id         int
	Customer   CustomerReal
	Item       string
	Restaurant string
	Status     string
}

func RestaurantWorkerReal(ctx context.Context, id int, orderChan <-chan OrderReal, cookedChan chan<- OrderReal, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Restaurant Worker %d: Shutting Down \n", id)
			return
		case order, ok := <-orderChan:
			if !ok {
				return
			}
			fmt.Printf("Restaurant Worker %d: Cooking %s for %s \n", id, order.Item, order.Customer.Name)
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

			order.Status = "Cooked"
			cookedChan <- order
		}
	}
}

func driverWorkerReal(ctx context.Context, id int, cookedChan <-chan OrderReal, deliveredChan chan<- OrderReal, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Driver %d: Shutting Down \n", id)
			return
		case order, ok := <-cookedChan:
			if !ok {
				return
			}
			fmt.Printf("Driver %d: Delivering %s to %s \n", id, order.Item, order.Customer.Name)
			time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second)

			order.Status = "Delivered"
			deliveredChan <- order
		}
	}
}

func dispatcherReal(ctx context.Context, orders []OrderReal, orderChan chan<- OrderReal) {
	for _, o := range orders {
		select {
		case <-ctx.Done():
			fmt.Println("Dispatcher Cancelled")
			return
		case orderChan <- o:
			fmt.Printf("Dispatcher: Sent order %d (%s) from %s \n", o.Id, o.Item, o.Customer.Name)
		}
	}
	close(orderChan)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	customers := []CustomerReal{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
		{4, "Diana"},
		{5, "Eve"},
	}

	orders := []OrderReal{
		{1, customers[0], "Burger", "FastFood", "New"},
		{2, customers[1], "Pizza", "FastFood", "New"},
		{3, customers[2], "Sushi", "Japanese", "New"},
		{4, customers[3], "Ramen", "Japanese", "New"},
		{5, customers[4], "Steak", "Western", "New"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	orderChan := make(chan OrderReal, len(orders))
	cookedChan := make(chan OrderReal, len(orders))
	deliveredChan := make(chan OrderReal, len(orders))

	var wgRestaurant, wgDriver sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wgRestaurant.Add(1)
		go RestaurantWorkerReal(ctx, i, orderChan, cookedChan, &wgRestaurant)
	}

	for i := 1; i <= 2; i++ {
		wgDriver.Add(1)
		go driverWorkerReal(ctx, i, cookedChan, deliveredChan, &wgDriver)
	}

	go dispatcherReal(ctx, orders, orderChan)

	go func() {
		wgRestaurant.Wait()
		close(cookedChan)
	}()

	go func() {
		wgDriver.Wait()
		close(deliveredChan)
	}()

	for delivered := range deliveredChan {
		fmt.Printf("SYSTEM: Order %d for %s is %s \n", delivered.Id, delivered.Customer.Name, delivered.Status)
	}

	fmt.Println("SYSTEM: All orders processed or timeout reached")
}
