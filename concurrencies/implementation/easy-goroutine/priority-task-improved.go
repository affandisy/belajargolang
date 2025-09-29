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

type TaskImproved struct {
	Name     string
	Priority int
}

type DriverImproved struct {
	Id   int
	Name string
}

type Result struct {
	Priority int
	Message  string
}

func PriorityDispatcherImproved(tasks []TaskImproved, taskChan chan<- TaskImproved) {
	for priority := 1; priority <= 3; priority++ {
		for _, task := range tasks {
			if task.Priority == priority {
				taskChan <- task
			}
		}
	}
	close(taskChan)
}

func (d DriverImproved) assignTask(taskChan <-chan TaskImproved, wg *sync.WaitGroup, resultChan chan<- Result) {
	defer wg.Done()

	for task := range taskChan {
		time.Sleep(2 * time.Second)
		// fmt.Printf("Driver %s (%d): Processing Task: %s \n", d.Name, d.Id, task.Name)
		// resultChan <- fmt.Sprintf("Driver %s (%d): Processing Task: %s", d.Name, d.Id, task.Name)
		resultChan <- Result{
			Priority: task.Priority,
			Message:  fmt.Sprintf("Driver %s (%d): Processing Task: %s", d.Name, d.Id, task.Name),
		}
	}
}

func main() {
	tasks := []TaskImproved{
		{"Emergency Medical Transport", 1},
		{"Delivery at Zone A", 2},
		{"Pickup at Zone B", 3},
	}

	taskChan := make(chan TaskImproved)
	resultChan := make(chan Result, len(tasks))

	var wg sync.WaitGroup

	drivers := []DriverImproved{
		{1, "Syihabuddin Affandi"},
		{2, "Momonga King"},
		{3, "Kazuya Sinho"},
		{4, "Rayleigh Silver"},
	}

	for _, driver := range drivers {
		wg.Add(1)
		go driver.assignTask(taskChan, &wg, resultChan)
	}

	go PriorityDispatcherImproved(tasks, taskChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	results := make(map[int]string)

	for res := range resultChan {
		results[res.Priority] = res.Message
	}

	for i := 1; i <= len(tasks); i++ {
		if msg, ok := results[i]; ok {
			fmt.Println(msg)
		}
	}

	// wg.Wait()
	fmt.Println("All Tasks are completed")
}
