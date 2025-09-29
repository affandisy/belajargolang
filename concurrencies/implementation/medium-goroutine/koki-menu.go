package main

import (
	"fmt"
	"sync"
	"time"
)

// Algoritma Priority Goroutines
// Definisikan Data:
// Buat Struct Order dengan Nama & Priorities
// Buat Struct Chefs dengan Id & Nama
// Buat Struct ResultRes untuk hasil kerja
// Menyiapkan data Awal:
// Menggunakan Slice untuk mengisi Struct Order dengan pesanan & priorities
// Menggunakan Slice untuk mengisi Struct Chefs dengan Id & Nama
// Membuat Channel:
// Channel orderChan untuk mengirim order ke Chef
// Channel resultChan untuk mengumpulkan hasil kerja
// Menjalankan goroutine chef:
// Looping setiap Chef
// Menjalankan fungsi assignTaskRestaurant dalam goroutine
// Tambahkan counter WaitGroup
// Jalankan dispatcher:
// Fungsi PriorityDispatcherRestaurant mengirim order ke channel sesuai urutan prioritas
// Tutup channel setelah selesai
// Tutup resultChan setelah semua goroutine selesai
// Goroutine anonim menunggu wg.Wait(), lalu menutup resultChan
// Kumpulkan hasil:
// Loop membaca dari resultChan
// Simpan hasil ke map dengan key = prioritas
// Cetak hasil sesuai urutan prioritas:
// Looping dari 1 sampai jumlah orders
// Cetak hasil jika ada

type MenuPriority struct {
	Menu     string
	Priority int
}

type Koki struct {
	Name string
	Id   int
}

type ResultNew struct {
	Priority int
	Message  string
}

func (k Koki) assignTaskRestaurant(menuChan <-chan MenuPriority, wg *sync.WaitGroup, resultChan chan<- ResultNew) {
	defer wg.Done()

	for menu := range menuChan {
		time.Sleep(3 * time.Second) // Simulasi Memasak
		resultChan <- ResultNew{
			Priority: menu.Priority,
			Message:  fmt.Sprintf("Koki %s dengan ID Tugas %d: Sedang Memasak Menu MBG: %s", k.Name, k.Id, menu.Menu),
		}
	}
}

func PriorityDispatcherMBG(menus []MenuPriority, menuChan chan<- MenuPriority) {
	for priority := 1; priority <= 3; priority++ {
		for _, menu := range menus {
			if menu.Priority == priority {
				menuChan <- menu
			}
		}
	}
	close(menuChan)
}

func main() {
	// Mengisi data
	menus := []MenuPriority{
		{"Nasi Goreng", 1},
		{"Kwetiau", 1},
		{"Capcay", 2},
		{"Mie Ayam", 2},
		{"Salad", 3},
		{"Gorengan", 3},
	}

	kokis := []Koki{
		{"Syihabuddin Affandi", 1},
		{"Dzikri Ramadhan", 2},
		{"Momonga Kingslayer", 3},
		{"Jhon Koo Wee", 4},
		{"Gibs Gajah", 5},
	}

	for _, value := range menus {
		fmt.Printf("Menu di MBG adalah %s dengan Prioritas %d \n", value.Menu, value.Priority)
	}

	for _, value := range kokis {
		fmt.Printf("Koki di MBG adalah %s dengan ID Tugas: %d \n", value.Name, value.Id)
	}

	// Variabel wg untuk WaitGroup
	var wg sync.WaitGroup

	// Membuat Channel Order dan Result
	menuChan := make(chan MenuPriority)
	resultChan := make(chan ResultNew, len(menus))

	// Menjalankan goroutine pekerja koki
	for _, koki := range kokis {
		wg.Add(1)
		go koki.assignTaskRestaurant(menuChan, &wg, resultChan)
	}

	// Menjalankan Priority Dispatcher
	PriorityDispatcherMBG(menus, menuChan)

	// Menutup channel resultChan
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Membuat slice untuk menampung result dari channel resultChan
	results := make(map[int][]string)

	for res := range resultChan {
		results[res.Priority] = append(results[res.Priority], res.Message)
	}

	// Cetak sesuai dengan prioritas
	for i := 1; i <= 3; i++ {
		if msgs, ok := results[i]; ok {
			for _, msg := range msgs {
				fmt.Println(msg)
			}
		}
	}

	// Task Completed
	fmt.Println("All Menus Are Completed")

}
