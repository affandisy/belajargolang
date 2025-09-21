package main

import (
	"fmt"
	"sort"
	"strings"
)

func hitungStatistik(label string, data []float64) {
	fmt.Printf("Analisa %s \n", label)

	if len(data) == 0 {
		fmt.Printf("Slice kosong \n")
		return
	}

	var jumlah float64
	for _, nilai := range data {
		jumlah = jumlah + nilai
	}

	fmt.Printf("Jumlah semua angka adalah %2.f \n", jumlah)

	// Rata-rata
	rataRata := jumlah / float64(len(data))
	fmt.Printf("Rata-rata dari slice adalah %2.f \n", rataRata)

	// Median
	dataSorted := make([]float64, len(data))
	copy(dataSorted, data)

	sort.Float64s(dataSorted)

	var median float64
	if len(dataSorted)%2 != 0 {
		median = dataSorted[len(dataSorted)/2]
	} else {
		tengah1 := dataSorted[len(dataSorted)/2-1]
		tengah2 := dataSorted[len(dataSorted)/2]
		median = (tengah1 + tengah2) / 2.0
	}
	fmt.Printf("Median dari slice adalah: %2.f \n", median)

	fmt.Println()
}

func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func checkXOXO(str string) bool {
	str = strings.ToLower(str)

	var xCount int
	var oCount int

	for _, char := range str {
		switch char {
		case 'x':
			xCount++
		case 'o':
			oCount++
		}
	}

	return xCount == oCount
}

func AlayGen(kata ...string) string {
	aturanAlay := map[rune]rune{
		'a': '4',
		'e': '3',
		'i': '!',
		'l': '1',
		'n': 'N',
		's': '5',
		'x': '*',
	}

	var hasil strings.Builder

	for _, k := range kata {
		for _, karakter := range k {
			if pengganti, ok := aturanAlay[karakter]; ok {
				hasil.WriteRune(pengganti)
			} else {
				hasil.WriteRune(karakter)
			}
		}
		hasil.WriteRune(' ')
	}
	return strings.TrimSpace(hasil.String())

}

func fibonacciIteratif(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	var hasil int

	for i := 2; i <= n; i++ {
		hasil = a + b
		a = b
		b = hasil
	}

	return hasil
}

func main() {
	// 	Buatlah sebuah `variable` yang berisikan `slice` dari `float64`, dan hitunglah `rata-rata`, `jumlah`, dan `median` dari slice tersebut.
	// Gunakan for dan operasi aritmatika.

	// go
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}
	kata1 := "katak"
	kata2 := "hello"
	kata3 := "level"
	s1 := "xoxoxoxoxoxoxoxo"
	s2 := "xoxo"
	s3 := "xoxoxoxoxoxoxo"
	angkaFibonacci := 10

	hitungStatistik("Slice 1", slice1)
	hitungStatistik("Slice 2", slice2)

	fmt.Printf("%s adalah Palindrome: %t \n", kata1, isPalindrome(kata1))
	fmt.Printf("%s adalah Palindrome: %t \n", kata2, isPalindrome(kata2))
	fmt.Printf("%s adalah Palindrome: %t \n", kata3, isPalindrome(kata3))

	fmt.Println("---------------")

	fmt.Printf("%s memiliki jumlah 'X' dan 'O' yang sama? %t \n", s1, checkXOXO(s1))
	fmt.Printf("%s memiliki jumlah 'X' dan 'O' yang sama? %t \n", s2, checkXOXO(s2))
	fmt.Printf("%s memiliki jumlah 'X' dan 'O' yang sama? %t \n", s3, checkXOXO(s3))

	// var kata string
	slice3 := []int{1, 3, 2, 5, 9, 20, 11, 12, 15, 11, 15, 30, 7}

	for i := 0; i < len(slice3); i++ {
		for j := 0; j < len(slice3)-1-i; j++ {
			if slice3[j] > slice3[j+1] {
				slice3[j], slice3[j+1] = slice3[j+1], slice3[j]
			}
		}
	}

	fmt.Println("Slice setelah diurutkan: ", slice3)

	fmt.Println(AlayGen("helo", "dunia", "ini", "kata", "alay", "dari", "tahun", "awal", "duaribuan"))
	fmt.Println(AlayGen("ini", "adalah", "contoh", "kalimat", "alay"))
	fmt.Println(AlayGen("sukses"))

	fmt.Printf("Deret Fibonacci ke-%d adalah: %d \n", angkaFibonacci, fibonacciIteratif(angkaFibonacci))

}
