package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Buatlah sebuah variabel dengan nama `myNum` dengan tipe data int32 dan assign nilai 50 ke dalam variabel tersebut Lalu, print variable tersebut ke dalam terminal
	var myNum int32 = 50
	fmt.Println(myNum)
	// Buatlah sebuah variabel dengan nama `myNum2` dengan tipe data float32 dan assign nilai 51 ke dalam variabel tersebut Lalu, print variable tersebut ke dalam terminal
	var myNum2 float32 = 51.2
	fmt.Println(myNum2)
	// Buatlah sebuah variabel dengan nama `myNumStr` dengan tipe data string dan assign karakter "50" ke dalam variabel tersebut Lalu, print variable tersebut ke dalam terminal
	var myNumStr string = "50"
	fmt.Println(myNumStr)

	// Buatlah variabel dengan dengan ketentuan berikut- x dan assign nilai int32 5- y dan assign nilai int32 10
	// Buatlah variabel baru dengan nama z, yang merupakan hasil penjumlahan dari x dan y.
	// Lalu print ke dalam terminal
	var x int32 = 5
	var y int32 = 10
	var z int32
	z = x + y
	fmt.Println(z)

	// Membuat program CLI yang menerima input nama (string) dan akan menampilkan `hello [nama]`
	// contoh ketika kita memasukan input `Bambang` maka output yang di print adalah `hello Bambang"

	reader := bufio.NewReader(os.Stdin)

	// fmt.Println("Masukkan nama lengkap anda: ")

	namaLengkap, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error saat Membaca Input: ", err)
		return
	}

	namaLengkap = strings.TrimSpace(namaLengkap)

	// fmt.Printf("Hello %s \n", namaLengkap)

	// Buatlah slice dengan nama `people` yang berisikan nama-nama orang `Walt, Jesse, Skyler, Saul`-
	// hitung berapa panjang slice tersebut dan tampilkan ke dalam terminal-
	// tambahkan `Hank` dan `Marie` ke dalam slice `people`-
	// hitung berapa panjang slice setelah ditambahkan `Hank` dan `Marie` dan tampilkan ke dalam terminal-
	// tampilkan `people` ke dalam terminal

	people := []string{"Walt", "Jesse", "Skyler", "Saul"}

	fmt.Println(len(people))

	people = append(people, "Hank", "Marie")

	fmt.Println(len(people))

	fmt.Println("Menampilkan People: ", people)

	// Buatlah sebuah array dengan kapasitas 3 elemen yang berisikan `map` yang memiliki `keys`, `name` dan `gender`, lalu isi array tersebut dengan data berikut, menggunakan perintah append
	// name: Hank
	// gender: M
	// name: Heisenberg
	// gender: M
	// name: Skyler
	// gender: F
	// dan tampilkan array tersebut ke dalam terminal.
	// Setelah itu modifikasi array tadi, untuk setiap data dengan `gender` `F` tambahkan `Mrs` di depan namanya, dan `M` tambahkan `Mr`.
	// Lalu tampilkan kembali ke dalam terminal

	users := make([]map[string]string, 0, 3)

	users = append(users, map[string]string{"name": "Hank", "gender": "M"})
	users = append(users, map[string]string{"name": "Heisenberg", "gender": "M"})
	users = append(users, map[string]string{"name": "Skyler", "gender": "F"})

	fmt.Println("Data Awal: ", users)

	for _, user := range users {
		if user["gender"] == "F" {
			user["name"] = "Mrs. " + user["name"]
		} else if user["gender"] == "M" {
			user["name"] = "Mr. " + user["name"]
		}
	}

	fmt.Println("Data setelah dimodifikasi: ", users)

}
