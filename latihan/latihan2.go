package main

import (
	"fmt"
	"strings"
)

func checkSlice(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	// Filter Slice Ganjil
	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sliceKosong := []int{}

	for _, slice := range slice3 {
		if slice%2 != 0 {
			sliceKosong = append(sliceKosong, slice)
		}
	}

	fmt.Println(sliceKosong)

	// String Reversal
	var inputString1 string

	fmt.Print("Input string anda \n")
	fmt.Scan(&inputString1)

	runes := []rune(inputString1)

	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - 1 - i

		runes[i], runes[j] = runes[j], runes[i]
	}

	reversedStrings := string(runes)

	fmt.Println(inputString1)
	fmt.Println(reversedStrings)

	// Mengecek duplikat di slice
	sliceDuplikat := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceDuplikat1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}

	fmt.Printf("Slice ini %v ada duplikat?: %t \n", sliceDuplikat, checkSlice(sliceDuplikat))
	fmt.Printf("Slice ini %v ada duplikat?: %t \n", sliceDuplikat1, checkSlice(sliceDuplikat1))

	// Menghitung setiap kata yang keluar
	kalimat := "saya suka makan nasi dan saya suka minum kopi setiap hari"

	kataKata := strings.Fields(kalimat)

	hitunganKata := make(map[string]int)

	for _, kata := range kataKata {
		hitunganKata[kata] = hitunganKata[kata] + 1
	}

	fmt.Println("Hasil dari hitungan kata: ")
	for kata, jumlah := range hitunganKata {
		fmt.Printf(" - %s telah muncul sebanyak %d \n", kata, jumlah)
	}

	// Kalkulator Faktorial
	var penampungInput int

	fmt.Print("Tolong masukkan angka anda: ")
	fmt.Scan(&penampungInput)

	var penampung = 1

	for i := 1; i <= penampungInput; i++ {
		penampung = penampung * i
	}

	fmt.Printf("Hasil dari faktorial dari %d adalah %d \n", penampungInput, penampung)

	// Check apakah Palindrome

	kalimat1 := "A man a plan a canal Panama"

	cleanedKalimat := strings.ToLower(strings.ReplaceAll(kalimat1, " ", ""))

	runes1 := []rune(cleanedKalimat)

	left := 0
	right := len(runes1) - 1

	isPalindrome := true

	for left < right {
		if runes1[left] != runes1[right] {
			isPalindrome = false
			break
		}

		left++
		right--
	}

	if isPalindrome {
		fmt.Printf("Kalimat %s adalah sebuah palindrome \n", kalimat1)
	} else {
		fmt.Printf("Kalimat %s adalah bukan palindrome \n", kalimat1)
	}

	// Rotasi slice
	numberSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	firstElement := numberSlice[0]

	numberSlice = numberSlice[1:]

	numberSlice = append(numberSlice, firstElement)

	fmt.Println("Number slice setelah rotasi: ", numberSlice)

	// Bintang Segitiga
	var pembatas int = 5

	for i := 1; i <= pembatas; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	for i := pembatas; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// Check Password
	password := "pass"
	angkaValid := "123456789"

	panjangCukup := false
	adaAngka := false

	if len(password) >= 8 {
		panjangCukup = true
	}

	for _, karakterPassword := range password {
		for _, karakterAngka := range angkaValid {
			if karakterPassword == karakterAngka {
				adaAngka = true
				break
			}
		}
		if adaAngka {
			break
		}
	}

	if panjangCukup && adaAngka {
		fmt.Println("Password Valid")
	} else {
		fmt.Println("Tidak valid")
		if !panjangCukup {
			fmt.Println("Kurang panjang")
		}
		if !adaAngka {
			fmt.Println("Kurang angka")
		}
	}

}
