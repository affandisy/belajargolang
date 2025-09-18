package main

import (
	"fmt"
)

func main() {
	// Variabel yang menyimpan alamat dari variabel lain
	nama := "Syihabuddin Affandi"
	namaRef := &nama // Alamat Memory

	fmt.Println("Nama saya adalah: ", nama)
	fmt.Println("Nama referensi adalah: ", namaRef)
	fmt.Println("Value Referensi adalah: ", *namaRef)

	*namaRef = "Muhammad Firmansyah"

	fmt.Println("Nama Ref setelah diubah: ", *namaRef)
	fmt.Println("Nama setelah diubah: ", nama)

	namaBaru := "Syihabuddin Affandi"

	fmt.Println("Value from variable: ", namaBaru)
	fmt.Println("Alamat Memory: ", &namaBaru)

	angkaAsli := 50

	// Tanpa Pointer
	nilaiPointer := ubahNilaiTanpaPointer(angkaAsli)
	fmt.Println("TANPA pointer: ", nilaiPointer)

	// Dengan Pointer
	ubahNilaiDenganPointer(&angkaAsli)
	fmt.Println("DENGAN pointer: ", angkaAsli)

}

func ubahNilaiDenganPointer(nilai *int) {
	*nilai = *nilai + 10
}

func ubahNilaiTanpaPointer(nilai int) int {
	return nilai + 10
}
