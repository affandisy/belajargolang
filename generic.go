package main

import "fmt"

type Numbers interface {
	int | float64 | float32
}

func SumInterface[T Numbers](nums []T) T {
	var total T
	for _, v := range nums {
		total = total + v
	}

	return total
}

func sumInt(nums []int) (total int) {
	for _, n := range nums {
		total = total + n
	}

	return total
}

func sumFloat(nums []float64) (total float64) {
	for _, n := range nums {
		total = total + n
	}

	return total
}

func Sum[T int | float64](nums []T) T {
	var total T
	for _, v := range nums {
		total = total + v
	}

	return total
}

// Type data any kita bisa memasukkan data apapun, any aliasing dari interface{}
func PrintAny[T any](val T) {
	fmt.Println(val)
}

// Meskipun menggunakan go generic, tapi di dalam golangnya kita tidak bisa meng-compare dua tipe data yang berbeda
func IsEqual[T comparable](a, b T) bool {
	return a == b
}

type BoxFlex[T any] struct {
	Value T
}

type Box struct {
	Value int
}

type BoxInterface struct {
	Value interface{}
}

func main() {
	paramInt := []int{10, 30, 20, 120}
	paramFloat := []float64{12.6, 13.8, 15.2}

	sumInt(paramInt)
	sumFloat(paramFloat)

	fmt.Println("Float Go Generic: ", Sum(paramFloat))
	fmt.Println("Integer Go Generic: ", Sum(paramInt))

	fmt.Println("Float Go Generic: ", SumInterface(paramFloat))
	fmt.Println("Integer Go Generic: ", SumInterface(paramInt))

	PrintAny("Syihabuddin Affandi")
	PrintAny(122)
	PrintAny(true)
	PrintAny(map[string]interface{}{
		"Data": 1,
	})

	fmt.Println(IsEqual(10, 12))
	fmt.Println(IsEqual("Gooo", "Laaang"))

	intBox := BoxFlex[int]{
		Value: 10,
	}

	fmt.Println(intBox)

	intBox1 := Box{
		Value: 10,
	}

	fmt.Println(intBox1)

	intBox2 := BoxInterface{
		Value: 10.10,
	}

	fmt.Println(intBox2)
}
