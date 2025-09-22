package main

import "fmt"

func tambah(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Testing Ini")

	var var1 string = "This is string"
	var int1 int = 10000
	var float1 float64 = 12.13
	var bool1 bool = true
	slice1 := []int{1, 2, 3, 8, 4, 5, 6, 9, 11, 10, 13}

	fmt.Println(slice1)
	fmt.Println(var1)
	fmt.Println(int1)
	fmt.Println(float1)
	fmt.Println(bool1)

	if int1 < 100 {
		fmt.Println("Kurang dari 1000")
	} else {
		fmt.Println("Lebih dari sama dengan 1000")
	}

	for i := 0; i < 3; i++ {
		fmt.Println(slice1)
	}

	maxNumber := slice1[0]

	// Untuk mencari nilai tertinggi dari slice1
	for i := 1; i < len(slice1); i++ {
		if slice1[i] > maxNumber {
			maxNumber = slice1[i]
		}
	}

	fmt.Printf("Nilai tertinggi di slice1 adalah %d \n", maxNumber)

	nilaiDicari := 5
	ditemukan := false

	// Mencari nilai yang sama
	for _, item := range slice1 {
		if item == nilaiDicari {
			ditemukan = true
			break
		}
	}

	if ditemukan {
		fmt.Printf("Nilai anda adalah %d ditemukan %t \n", nilaiDicari, ditemukan)
	} else {
		fmt.Printf("Nilai anda tidak ditemukan")
	}

	// Menyapa pengguna
	var namaPengguna string
	var usiaPengguna int

	fmt.Print("Masukkan nama anda: ")
	fmt.Scan(&namaPengguna)

	fmt.Print("Masukkan usia anda: ")
	fmt.Scan(&usiaPengguna)

	fmt.Printf("Halo, Nama saya adalah %s dengan usia %d \n", namaPengguna, usiaPengguna)

	// Kalkulator Persegi panjang
	var lebarPersegiPanjang int
	var panjangPersegiPanjang int

	fmt.Print("Masukkan Panjang Persegi Panjang: ")
	fmt.Scan(&panjangPersegiPanjang)

	fmt.Print("Masukkan Lebar Persegi Panjang: ")
	fmt.Scan(&lebarPersegiPanjang)

	luasPersegiPanjang := panjangPersegiPanjang * lebarPersegiPanjang

	fmt.Printf("Luas persegi panjang adalah: %d \n", luasPersegiPanjang)

	// Apakah input tersebut positif, negatif, atau nol

	var angkaInput int

	fmt.Print("Masukkan angka anda: ")
	fmt.Scan(&angkaInput)

	if angkaInput == 0 {
		fmt.Println("Angka tersebut adalah 0")
	} else if angkaInput < 0 {
		fmt.Println("Angka tersebut negatif")
	} else {
		fmt.Println("Angka tersebut positif")
	}

	// Mencetak nilai 1 sampai N
	var angkaInput2 int

	fmt.Print("Masukkan angka anda: ")
	fmt.Scan(&angkaInput2)

	for i := 1; i <= angkaInput2; i++ {
		fmt.Println(i)
	}

	// Mencari bilangan genap
	var angkaInput3 int

	fmt.Print("Masukkan angka anda: ")
	fmt.Scan(&angkaInput3)

	for i := 0; i <= angkaInput3; i++ {
		if i%2 == 0 {
			var hasilGenap = i
			fmt.Println(hasilGenap)
		}
	}

	// Mencari jumlah string
	var inputString string

	fmt.Print("Masukkan string anda: ")
	fmt.Scan(&inputString)

	var jumlahString = len(inputString)

	fmt.Println(jumlahString)

	// Jumlah Elemen Slice
	slice2 := []int{1, 2, 3, 8, 4, 5, 6, 9, 11, 10, 13}

	var totalVariabel int

	for _, total := range slice2 {
		totalVariabel = totalVariabel + total
	}

	fmt.Println(totalVariabel)

	// FizzBuzz Klasik
	var lowerBound int
	var upperBound int

	fmt.Print("Masukkan lower bound: ")
	fmt.Scan(&lowerBound)

	fmt.Print("Masukkan upper bound: ")
	fmt.Scan(&upperBound)

	for i := lowerBound; i <= upperBound; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// Fungsi Penjumlahan
	var varTambah = tambah(10, 20)
	fmt.Println(varTambah)

	// Konversi Celcius ke Fahrenheit
	var inputCelcius float64

	fmt.Print("Masukkan input celcius: ")
	fmt.Scan(&inputCelcius)

	var fahrenheit float64 = (inputCelcius * 9.0 / 5.0) + 32

	fmt.Println(fahrenheit)
}
