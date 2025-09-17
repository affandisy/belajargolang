package main

import (
	"fmt"
)

func main() {
	buah := []string{"Nanas", "Apple", "Strawberry", "Naga"}

	fmt.Println(buah)

	buah = append(buah, "Jeruk")

	fmt.Println("Length Buah Slice: ", len(buah))
	fmt.Println("Capacity Buah Slice: ", cap(buah))

	buah2 := [5]string{"Nanas", "Apple", "Strawberry"}
	buah2[3] = "Salak"

	fmt.Println("Length Buah Array: ", len(buah2))
	fmt.Println("Capacity Buah Array: ", cap(buah2))

	fmt.Println("Buah Slice Pertama", buah)
	fmt.Println("Buah Array Pertama", buah2)

	buah3 := make([]string, 0, 10)
	buah3 = append(buah3, "Nanas", "Apple")

	fmt.Println("Buah Array Otomatis", buah3)
	fmt.Println("Length Buah Array: ", len(buah3))
	fmt.Println("Capacity Buah Array: ", cap(buah3))

	buah4 := make([]string, 0, len(buah2))

	for _, dtBuah := range buah2 {
		buah4 = append(buah4, dtBuah)
	}

	copy(buah4, buah)

	fmt.Println(buah4)
	fmt.Println("Length Buah Array: ", len(buah4))
	fmt.Println("Capacity Buah Array: ", cap(buah4))
}
