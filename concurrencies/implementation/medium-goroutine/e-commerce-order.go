package main

import (
	"fmt"
	"sync"
	"time"
)

// 3 Tahap Pemrosesan Pesanan

// Mewakili sebuah pesanan pelanggan, bentuknya struct dengan NameItem dan Id
type OrderCommerce struct {
	NameItem string
	Id       int
}

// Producer membuat order dari slice orders, mengirim setiap order ke channel ordersChan, setelah selesai order ditutup
func orderProducer(orders []OrderCommerce, ordersChan chan<- OrderCommerce) {
	for _, order := range orders {
		fmt.Printf("Producing order: %s \n", order.NameItem)
		ordersChan <- order
	}
	close(ordersChan)
}

// Membaca order dari channel ordersChan, simulasi pembayaran 2 detik, setelah selesai order dikirim ke channel paidChan, wg.Done() menandakan bahwa worker selesai
func paymentWorker(id int, ordersChan <-chan OrderCommerce, paidChan chan<- OrderCommerce, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range ordersChan {
		fmt.Printf("PaymentWorker %d: Processing Payment for %s \n", id, order.NameItem)
		time.Sleep(2 * time.Second)
		paidChan <- order
	}
}

// Membaca order dari paidChan, simulasi pengiriman 1 detik, mencetak pesan bahwa order dikirim
func shippingWorker(id int, paidChan <-chan OrderCommerce, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range paidChan {
		fmt.Printf("ShippingWorker %d: Shipping %s \n", id, order.NameItem)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	orders := []OrderCommerce{
		{"Laptop", 1},
		{"Printer", 2},
		{"Keyboard Gaming", 3},
		{"USB Type-C", 4},
		{"Mouse Gaming", 5},
		{"Monitor", 6},
		{"Smartphone", 7},
	}

	// Dua channel pipeline
	// ordersChan untuk input dari producer -> payment
	// paidChan untuk input dari payment -> shipping
	// Buffer dipakai agar tidak terlalu sering blocking
	ordersChan := make(chan OrderCommerce, len(orders))
	paidChan := make(chan OrderCommerce, len(orders))

	// Dua WaitGroup: satu untuk semua payment worker, satu untuk semua shipping worker
	var wgPayment sync.WaitGroup
	var wgShipping sync.WaitGroup

	// Menjalankan pipeline producer
	go orderProducer(orders, ordersChan)

	// Menjalankan beberapa worker pembayaran secara paralel
	numPaymentWorkers := 2
	for i := 1; i <= numPaymentWorkers; i++ {
		wgPayment.Add(1)
		go paymentWorker(i, ordersChan, paidChan, &wgPayment)
	}

	// Menjalankan beberapa worker shipping secara paralel
	numShippingWorkers := 2
	for i := 1; i <= numShippingWorkers; i++ {
		wgShipping.Add(1)
		go shippingWorker(i, paidChan, &wgShipping)
	}

	// Tutup channel setelah stage 2 Selesai, setelah payment worker selesai, paidChan ditutup,
	// memberikan sinyal ke shipping worker tidak ada data baru lagi
	go func() {
		wgPayment.Wait()
		close(paidChan)
	}()

	wgShipping.Wait()

	fmt.Println("All orders processed and shipped")

}
