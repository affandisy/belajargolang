package main

import (
	"fmt"
	"sync"
	"time"
)

// Alur kerja Priority Task
// Konsep Dasar: Dispatcher -> Channel -> Driver
// Dispatcher = Pengirim tugas. Ia memilih urutan tugas berdasarkan prioritas dan mengirimkannya ke "channel"
// Channel = Kotak antrian (queue) yang bisa diisi oleh satu pihak (dispatcher) dan dibaca oleh pihak lain (driver)
// Driver = Pekerja. Mereka mengambil tugas dari "channel" dan mengeksekusinya

// Alur kerja: Task(array) -> Dispatcher -> Channel -> Driver (goroutine)

// Peran WaitGroup
// wg adalah alat sinkronisasi
// Setiap kali menjalankan goroutine driver.assignTask, menambahkan wg.Add(1)
// Di dalam fungsi goroutine, defer wg.Done() agar ketika goroutine selesai, counter WaitGroup berkurang
// wg.Wait() memastikan program tidak berhenti sampai semua goroutine driver selesai bekerja

// Tahapan Eksekusi
// Main membuat channel "taskChan"
// Main menjalankan dua goroutine driver. Mereka siap "menunggu" task masuk ke "taskChan"
// Main juga menjalankan goroutine "PriorityDispatcher", yang mulai mengirimkan tugas ke "taskChan" sesuai urutan prioritas
// Driver goroutine berebut tugas yang masuk ke channel
// Misalnya Driver 1 dapat task pertama, Driver 2 dapat task kedua
// Channel ditutup setelah semua task selesai dikirim

type TaskImprove struct {
	Name     string
	Priority int
}

type DriverImprove struct {
	Id   int
	Name string
}

func (d DriverImprove) assignTask(taskChan <-chan TaskImprove, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskChan {
		fmt.Printf("Driver %s (%d): Processing Task: %s \n", d.Name, d.Id, task.Name)
		time.Sleep(2 * time.Second)
	}
}

func PriorityDispatcherImprove(tasks []TaskImprove, taskChan chan<- TaskImprove) {
	for priority := 1; priority <= 3; priority++ {
		for _, task := range tasks {
			if task.Priority == priority {
				taskChan <- task
			}
		}
	}
	close(taskChan)
}

func main() {
	tasks := []TaskImprove{
		{"Emergency Medical Transport", 1},
		{"Delivery at Zone A", 2},
		{"Pickup at Zone B", 3},
	}

	taskChan := make(chan TaskImprove)

	var wg sync.WaitGroup

	drivers := []DriverImprove{
		{1, "Syihabuddin Affandi"},
		{2, "Momonga King"},
	}

	for _, driver := range drivers {
		wg.Add(1)
		go driver.assignTask(taskChan, &wg)
	}

	go PriorityDispatcherImprove(tasks, taskChan)

	wg.Wait()
	fmt.Println("All Tasks are completed")
}
