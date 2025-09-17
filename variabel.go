package main

import "fmt"

func main() {
	a := 20
	b := 2
	nilaiUjian := 80
	lulusAbsensi := true

	// Ini Komen
	var lulusNilai bool = nilaiUjian > 75

	var name string = "Syihabuddin Affandi"

	var title, jumlah, desc = "Buku Novel Terluka", 2, "Sebuah buku tentang terluka"

	c, d := "sihab", "ganteng"

	fmt.Println(c, d)
	fmt.Println("======")

	c, d = d, c
	fmt.Println(c, d)

	fmt.Println(title)
	fmt.Println(jumlah)
	fmt.Println(desc)

	fmt.Println("name: \n", name)

	fmt.Printf("Lulus ujian? %t\n", lulusNilai)

	if lulusNilai && lulusAbsensi {
		fmt.Println("Selamat anda lulus")
	} else {
		fmt.Println("Tidak lulus")
	}

	hasilJumlah := a + b
	hasilKali := a * b
	hasilBagi := a / b
	fmt.Printf("Penjumlahan: %d + %d = %d\n", a, b, hasilJumlah)

	fmt.Printf("Perkalian: %d * %d = %d\n", a, b, hasilKali)

	fmt.Printf("Pembagian: %d / %d = %d\n", a, b, hasilBagi)

	fmt.Println("Halo, Dunia!")
	fmt.Println("Isi Dunia: ", 2)

	namaLengkap := "Syihabuddin Affandi"

	role := "Backend Engineer"

	ukuranSepatu := 40

	fmt.Printf("Ja, %s (%s) ukuran sepatu kamu %d \n", namaLengkap, role, ukuranSepatu)

	isActive := true

	fmt.Println(isActive)

	var myNum int32 = 50

	fmt.Println(myNum)

	var myNum2 float32 = 51.2

	fmt.Println(myNum2)

	var myNumStr string = "50"

	fmt.Println(myNumStr)

	var x int32 = 5
	var y int32 = 10
	z := x + y

	fmt.Println(z)

}
