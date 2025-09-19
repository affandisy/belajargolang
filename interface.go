package main

import (
	"fmt"
	"math"
)

type UserInterface struct {
	Name       string
	IsActive   bool
	DetailData interface{}
}

// Definisi Interface BangunRuang
type BangunRuang interface {
	Luas() float64
	Keliling() float64
}

// Struct Lingkaran
type Lingkaran struct {
	JariJari float64
}

// Struct Persegi Panjang
type PersegiPanjang struct {
	Panjang float64
	Lebar   float64
}

// Method Luas Persegi Panjang
func (p PersegiPanjang) Luas() float64 {
	return p.Panjang * p.Lebar
}

// Method Keliling Persegi Panjang
func (p PersegiPanjang) Keliling() float64 {
	return 2 * (p.Panjang + p.Lebar)
}

// Method Luas Lingkaran
func (l Lingkaran) Luas() float64 {
	return math.Pi * l.JariJari * l.JariJari
}

// Method Keliling Lingkaran
func (l Lingkaran) Keliling() float64 {
	return 2 * math.Pi * l.JariJari
}

func PrintBangunRuangInfo(s BangunRuang) {
	fmt.Println("Keliling: ", s.Keliling())
	fmt.Println("Luas: ", s.Luas())
}

func main() {
	// Empty Interface
	var nama interface{} = "Syihabuddin Affandi"

	fmt.Println("Nama before: ", nama)

	user := UserInterface{
		Name:     "Syihabuddin Affandi",
		IsActive: true,
		// DetailData: map[string]interface{}{
		// 	"kelas":      10,
		// 	"wali_murid": "orangtua",
		// },
		DetailData: []string{"huda", "apel"},
	}

	val, ok := user.DetailData.([]string)

	// val, ok := user.DetailData.(map[string]interface{})
	fmt.Println(val[1], ok)

	fmt.Println(user.DetailData)

	// Interface
	l := Lingkaran{
		JariJari: 10,
	}

	p := PersegiPanjang{
		Panjang: 10,
		Lebar:   20,
	}

	x := l.Luas()
	fmt.Println("Luas Lingkaran: ", x)

	y := p.Luas()
	fmt.Println("Luas Persegi Panjang: ", y)

	PrintBangunRuangInfo(l)
	PrintBangunRuangInfo(p)
}
