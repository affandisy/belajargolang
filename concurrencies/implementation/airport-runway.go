package main

import (
	"fmt"
	"sync"
	"time"
)

// Membuat Struct Dulu

type Plane struct {
	Name     string
	Priority int
}

type Runway struct {
	Name string
}

func (r Runway) assignRunwayPlane(planeChan <-chan Plane, wg *sync.WaitGroup, resultChan chan<- string) {
	defer wg.Done()

	for plane := range planeChan {
		time.Sleep(3 * time.Second)

		resultChan <- fmt.Sprintf("%s: Plane %s landed (Priority %d)", r.Name, plane.Name, plane.Priority)
	}
}

func RunwayDispatcher(planes []Plane, planeChan chan<- Plane) {
	for priority := 1; priority <= 3; priority++ {
		for _, plane := range planes {
			if plane.Priority == priority {
				planeChan <- plane
			}
		}
	}
	close(planeChan)
}

func main() {

	planes := []Plane{
		{"Garuda-123", 1},   // Emergency
		{"Singapore-77", 2}, // International
		{"AirAsia-44", 3},   // Domestic
		{"Lion-55", 3},      // Domestic
		{"Qatar-88", 2},     // International
	}

	runways := []Runway{
		{"Runway 1"},
		{"Runway 2"},
		{"Runway 3"},
	}

	for _, plane := range planes {
		fmt.Printf("Pesawat %s dalam posisi Prioritas %d \n", plane.Name, plane.Priority)
	}

	for _, runway := range runways {
		fmt.Printf("%s \n", runway.Name)
	}

	// Membuat Channel
	planeChan := make(chan Plane)
	resultChan := make(chan string)

	// Membuat WaitGroup
	var wg sync.WaitGroup

	// Menjalankan goroutine worker runway
	for _, runway := range runways {
		wg.Add(1)
		go runway.assignRunwayPlane(planeChan, &wg, resultChan)
	}

	// Runway Dispatcher
	go RunwayDispatcher(planes, planeChan)

	// Menutup channel resultChan
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for msg := range resultChan {
		fmt.Println(msg)
	}

	// Membuat slice untuk menampung result dari channel resultChan
	// results := make(map[int][]string)

	// for res := range resultChan {
	// 	results[res.Priority] = append(results[res.Priority], res.Message)
	// }

	// // Cetak sesuai dengan prioritas
	// for i := 1; i <= 3; i++ {
	// 	if msgs, ok := results[i]; ok {
	// 		for _, msg := range msgs {
	// 			fmt.Println(msg)
	// 		}
	// 	}
	// }

	fmt.Println("All planes have landed safely")

}
