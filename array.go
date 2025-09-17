package main

import (
	"fmt"
	"reflect"
)

func main() {
	buah := [3]string{"Nanas", "Strawberry", "Apple"}

	buahBanyak := [...]string{"Nanas", "Strawberry", "Apple", "Pisang", "Melon", "Naga"}

	buahBanyakInterface := [...]interface{}{"Nanas", 1, "Salak", "Jeruk", "Apel"}

	fmt.Println("Buah: ", buah)
	fmt.Println("Buah Banyak: ", buahBanyak)
	fmt.Println("Berapa banyak?", len(buahBanyak))

	for _, dataBuah := range buah {
		fmt.Println(dataBuah)
	}

	for key, dataBuah1 := range buahBanyak {
		fmt.Println(key, dataBuah1)
	}

	for key, dataBuah2 := range buahBanyakInterface {
		typeBuah := reflect.TypeOf(dataBuah2)
		fmt.Println(key, dataBuah2, typeBuah)
	}
}
