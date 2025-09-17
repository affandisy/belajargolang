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
}
