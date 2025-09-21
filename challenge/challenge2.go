package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	// 	Buatlah sebuah program CLI sederhana yang menerima input`Nama`
	// Program tersebut akan me-random angka `1-100` dan akan menampilkan :

	// `Selamat [Name], anda sangat beruntung` jika mendapat angka > 80
	// `Selamat [Name], anda beruntung` jika mendapat angka <= 80 dan > 60
	// `Mohon maaf [Name], anda kurang beruntung` jika mendapat angka <= 60 dan > 40
	// `Mohon maaf [Name], anda sial` jika mendapat angka <= 40

	// Gunakan switch-case dalam studi kasus diatas, untuk membuat random angka silahkan explore `Math.Rand`NG Challenge 1 : Conditional 1

	// Deklarasi variabel penampung
	var namaInput string

	fmt.Print("Masukkan nama anda \n")
	fmt.Scanln(&namaInput)

	// Deklarasi angka penampung
	angka := rand.Intn(101)

	fmt.Printf("%d adalah angka random \n", angka)

	switch {
	case angka >= 80:
		fmt.Printf("Mohon maaf %s, anda sangat beruntung \n", namaInput)
	case angka >= 60:
		fmt.Printf("Mohon maaf %s, anda beruntung \n", namaInput)
	case angka >= 40:
		fmt.Printf("Mohon maaf %s, anda kurang beruntung \n", namaInput)
	case angka < 40:
		fmt.Printf("Mohon maaf %s, anda sial \n", namaInput)
	default:
		fmt.Println("Angka kurang atau lebih \n")
	}

	// 	Buatlah sebuah aplikasi CLI yang akan menerima input berupa `Nama` dan `Umur`
	// Jika`Umur`lebih besar dari 18 maka aplikasi akan menampilkan `Silahkan masuk`
	// dan jika`Umur` lebih kecil sama dengan 18 maka akan menampilkan `Dilarang masuk, maksimal umur 19`
	// Dan perlu perlu juga menampilkan pesan error jika `Umur < 0`, `Umur > 100`,
	// Umur bukan merupakan angka,
	// PESAN ERRORdibebaskan selama masuk akal.

	// Deklarasi variabel Nama dan Umur
	var nama string
	var umur string

	fmt.Print("Masukkan nama anda")
	fmt.Scan(&nama)

	fmt.Print("Masukkan umur anda")
	fmt.Scan(&umur)

	var err error

	UmurBaru, err := strconv.Atoi(umur)

	if err != nil {
		fmt.Println("Error")
		return
	}

	if UmurBaru > 100 || UmurBaru < 0 {
		fmt.Println("Error lebih tinggi dan rendah")
		return
	}

	if UmurBaru <= 19 {
		fmt.Println("Silahkan masuk")
	} else {
		fmt.Println("Dilarang masuk, minimal 19 tahun")
	}

	// 	Buatlah sebuah variable yang merupakan slice yang berisi data map orang-orang dengan data sebagai berikut :

	// name: Hank
	// Age: 50
	// Job: Polisi

	// name: Heisenberg
	// Age: 52
	// Job: Ilmuwan

	// name: Skyler
	// Age: 48
	// Job: Akuntan

	// Loop kedalam slice tersebut dan tampilkan tampilkan data-data tiap orang dengan format tulisan :
	// ```Hi Perkenalkan, Nama saya [Nama], umur saya [umur], dan saya bekerja sebagai [pekerjaan]

	dataMap := []map[string]string{
		{"Name": "Hank", "Age": "50", "Job": "Polisi"},
		{"Name": "Heisenberg", "Age": "52", "Job": "Ilmuwan"},
		{"Name": "Skyler", "Age": "48", "Job": "Akuntan"},
	}

	for _, data := range dataMap {
		fmt.Printf("Hi Perkenalkan, Nama saya %s, umur saya %s, dan saya bekerja sebagai %s \n", data["Name"], data["Age"], data["Job"])
	}

}
