package main

import (
	"fmt"
	"sync"
	"time"
)

// Struct Customer memisahkan data pelanggan dari order
// Struct Order menyimpan Customer sebagai field
// Pipeline: Customer -> Restaurant -> Order
//  +-----------+         +----------------+        +---------------+
//  | Customer  | ----->  |  orderChan     | -----> | Restaurant(s) |
//  +-----------+         +----------------+        +---------------+
//        |                                                |
//        |                                                v
//        |                                       +----------------+
//        |                                       |  cookedChan    |
//        |                                       +----------------+
//        |                                                |
//        |                                                v
//        |                                       +---------------+
//        +-------------------------------------> | Driver(s)     |
//                                                +---------------+
// Customer membuat pesanan, dikirim melalui channel orderChan
// Restaurant Workers membaca pesanan dari channel orderChan, memasak, lalu mengirim makanan yang sudah siap ke channel cookedChan
// Driver Workers membaca dari channel cookedChan dan mengantar makanan ke customer
// Setelah semua, order selesai

type Customer struct {
	Name string
	Id   int
}

type OrderDelivery struct {
	Id       int
	Customer Customer
	Item     string
}

// Stage 1: Customer membuat pesanan
func CustomerOrder(order OrderDelivery, orderChan chan<- OrderDelivery) {
	fmt.Printf("Customer %s ordered %s \n", order.Customer.Name, order.Item)
	orderChan <- order
}

// Stage 2: Restaurant Worker
func RestaurantWorker(id int, orderChan <-chan OrderDelivery, cookedChan chan<- OrderDelivery, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range orderChan {
		fmt.Printf("Restaurant Worker %d: Cooking %s for %s \n", id, order.Item, order.Customer.Name)
		time.Sleep(2 * time.Second)
		cookedChan <- order
	}
}

// Stage 3: Driver Worker
func DriverWorker(id int, cookedChan <-chan OrderDelivery, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range cookedChan {
		fmt.Printf("Driver %d: Delivering %s to %s \n", id, order.Item, order.Customer.Name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	customers := []Customer{
		{"Syihabuddin Affandi", 1},
		{"Dzikri Ramadhan", 2},
		{"Momonga Kingslayer", 3},
		{"Jhon Koo Wee", 4},
		{"Gibs Gajah", 5},
	}

	orders := []OrderDelivery{
		{1, customers[0], "Burger"},
		{2, customers[1], "Pizza"},
		{3, customers[2], "Sushi"},
		{4, customers[3], "Noodles"},
		{5, customers[4], "Fried Rice"},
	}

	orderChan := make(chan OrderDelivery, len(orders))
	cookedChan := make(chan OrderDelivery, len(orders))

	var wgRestaurant sync.WaitGroup
	var wgDriver sync.WaitGroup

	for i := 1; i <= 2; i++ {
		wgRestaurant.Add(1)
		go RestaurantWorker(i, orderChan, cookedChan, &wgRestaurant)
	}

	for i := 1; i <= 2; i++ {
		wgDriver.Add(1)
		go DriverWorker(i, cookedChan, &wgDriver)
	}

	for _, o := range orders {
		go CustomerOrder(o, orderChan)
	}

	go func() {
		time.Sleep(1 * time.Second)
		close(orderChan)
	}()

	go func() {
		wgRestaurant.Wait()
		close(cookedChan)
	}()

	wgDriver.Wait()

	fmt.Println("All orders have been delivered")

}
