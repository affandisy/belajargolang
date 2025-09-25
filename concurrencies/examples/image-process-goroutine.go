package main

import (
	"fmt"
	"sync"
	"time"
)

type Image struct {
	ImageURL string
}

func ProcessImage(imageURL string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing image: %s \n", imageURL)
	time.Sleep(time.Second * 3)
	fmt.Printf("Image processing completed: %s \n", imageURL)
}

func main() {
	imgProcess := []Image{
		{ImageURL: "https://example.com/image1.jpg"},
		{ImageURL: "https://example.com/image2.jpg"},
		{ImageURL: "https://example.com/image3.jpg"},
		{ImageURL: "https://example.com/image4.jpg"},
	}

	var wg sync.WaitGroup
	wg.Add(len(imgProcess))

	for _, img := range imgProcess {
		go ProcessImage(img.ImageURL, &wg)
	}

	wg.Wait()

	fmt.Println("Image processing started, main application continue")

	time.Sleep(time.Second * 3)

	fmt.Println("All image processing is done")
}
