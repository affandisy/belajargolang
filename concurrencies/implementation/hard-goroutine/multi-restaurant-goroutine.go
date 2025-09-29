package main

import (
	"fmt"
	"sync"
	"time"
)

// Struct
// Customer -> menyimpan data customer (Id, Name)
// Order -> menyimpan pesanan (Id, Customer, Item, Restaurant)
// Restaurant -> punya Id, Name, dan channel khusus untuk order
// Driver -> Menyimpan data Driver
// Dispatcher
// Menerima slice orders
// Melihat "order.Restaurant" dan mengirim ke channel restoran yang cocok
// Kalau tidak ada restoran, cetak pesan error
// Setelah semua order dikirim, semua channel restoran ditutup agar worker tahu tidak ada order baru
// Restaurant Worker
// Menerima order dari channel khusus restoran
// Simulasi masak 2 detik
// Kirim order ke channel global cookedChan
// Driver Worker
// Menerima order dari cookedChan
// Simulasi antar 1 detik
// Cetak pesan bahwa order sudah diantar
// Main
// Buat data customers dan orders
// Buat restoran dalam "map[string]Restaurant" -> memudahkan routing order
// Jalankan goroutine untuk tiap restoran (masak paralel)
// Jalankan goroutine untuk driver (antar paralel)
// Jalankan dispatcher untuk mengirim order ke restoran yang sesuai
// Tutup cookedChan setelah semua restoran selesai
// Tunggu driver selesai, cetak pesan akhir

type CustomerRestaurant struct {
	Id   int
	Name string
}

type OrderRestaurant struct {
	Id         int
	Customer   CustomerRestaurant
	Item       string
	Restaurant string
}

type Restaurant struct {
	Id   int
	Name string
	Chan chan OrderRestaurant
}

type Driver struct {
	Id   int
	Name string
}

func dispatcher(orders []OrderRestaurant, restaurants map[string]Restaurant) {
	for _, order := range orders {
		if r, ok := restaurants[order.Restaurant]; ok {
			fmt.Printf("Dispatcher: Sending order %s (%s) to %s \n", order.Item, order.Customer.Name, r.Name)
			r.Chan <- order
		} else {
			fmt.Printf("Dispatcher: No Restaurant found for Order %s from %s \n", order.Item, order.Customer.Name)
		}
	}

	for _, r := range restaurants {
		close(r.Chan)
	}
}

func restaurantWorker(r Restaurant, cookedChan chan<- OrderRestaurant, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range r.Chan {
		fmt.Printf("%s: Cooking %s for %s \n", r.Name, order.Item, order.Customer.Name)

		time.Sleep(2 * time.Second)
		cookedChan <- order
	}
}

func driverWorker(d Driver, cookedChan <-chan OrderRestaurant, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range cookedChan {
		fmt.Printf("Driver %s: Delivering %s to %s \n", d.Name, order.Item, order.Customer.Name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	customers := []CustomerRestaurant{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
		{4, "Diana"},
		{5, "Eve"},
	}

	orders := []OrderRestaurant{
		{1, customers[0], "Burger", "FastFood"},
		{2, customers[1], "Pizza", "FastFood"},
		{3, customers[2], "Sushi", "Japanese"},
		{4, customers[3], "Ramen", "Japanese"},
		{5, customers[4], "Steak", "Western"},
	}

	restaurants := map[string]Restaurant{
		"FastFood": {1, "FastFood", make(chan OrderRestaurant)},
		"Japanese": {2, "Japanese", make(chan OrderRestaurant)},
		"Western":  {3, "Western", make(chan OrderRestaurant)},
	}

	cookedChan := make(chan OrderRestaurant)

	var wgRestaurant sync.WaitGroup
	var wgDriver sync.WaitGroup

	for _, r := range restaurants {
		wgRestaurant.Add(1)
		go restaurantWorker(r, cookedChan, &wgRestaurant)
	}

	drivers := []Driver{
		{1, "John"},
		{2, "Mike"},
		{3, "Joan"},
		{4, "John"},
		{5, "Doe"},
	}

	for _, d := range drivers {
		wgDriver.Add(1)
		go driverWorker(d, cookedChan, &wgDriver)
	}

	go dispatcher(orders, restaurants)

	go func() {
		wgRestaurant.Wait()
		close(cookedChan)
	}()

	wgDriver.Wait()

	fmt.Println("All orders have been delivered")

}
