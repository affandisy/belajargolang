package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Buku struct {
	Name  string
	Stock int
}

type Anggota struct {
	Name   string
	Pinjam []Pinjam
}

type Pinjam struct {
	Item     Buku
	Quantity int
}

func main() {
	fmt.Println("Perpustakaan Digital Hacktiv8 \n")

	daftarBuku := []Buku{
		{"Go Programming", 5},
		{"Data Structures", 3},
		{"Algorithms", 4},
		{"Database Systems", 2},
		{"Networking Basics", 6},
		{"Java Programming", 4},
	}

	for _, buku := range daftarBuku {
		fmt.Printf("%s - Stock %d \n", buku.Name, buku.Stock)
	}

	var anggotas []Anggota

	for {
		var penampungNama string

		fmt.Print("Nama anda: (Type 'done' if you done) \n")
		fmt.Scanln(&penampungNama)

		if penampungNama == "done" {
			break
		}

		anggota := Anggota{Name: penampungNama}

		fmt.Printf("Welcome %s ke perpustakaan digital \n", penampungNama)

		for {
			var penampungItem string
			var penampungQuantity int

			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Buku yang ingin dipinjam? (Type 'done' if you done) ")
			penampungItem, _ = reader.ReadString('\n')
			penampungItem = strings.TrimSpace(penampungItem)
			penampungItem = strings.ToLower(penampungItem)

			if penampungItem == "done" {
				break
			}

			// var chosen Buku
			found := false

			for i := range daftarBuku {
				if strings.ToLower(daftarBuku[i].Name) == penampungItem {
					found = true
					chosen := daftarBuku[i]

					fmt.Printf("Berapa banyak %s yang dibutuhkan? (Stock: %d)", chosen.Name, chosen.Stock)
					fmt.Scanln(&penampungQuantity)

					// Validasi input
					if penampungQuantity <= 0 {
						fmt.Println("Jumlah harus lebih dari 0 \n")
						break
					}

					if penampungQuantity > daftarBuku[i].Stock {
						fmt.Printf("Stock %s hanya tersedia %d, tidak bisa meminjam %d \n", chosen.Name, daftarBuku[i].Stock, penampungQuantity)
						break
					}

					// Menyimpan pinjaman
					anggota.Pinjam = append(anggota.Pinjam, Pinjam{Item: chosen, Quantity: penampungQuantity})

					// Mengurangi Stock buku
					daftarBuku[i].Stock = daftarBuku[i].Stock - penampungQuantity
					break
				}
			}

			if !found {
				fmt.Println("Buku tidak ada di daftar \n")
			}
		}

		anggotas = append(anggotas, anggota)
	}

	fmt.Println("\n ==== Loan Summary Per Member === \n")

	for _, a := range anggotas {
		fmt.Printf("%s borrowed: \n", a.Name)
		for _, p := range a.Pinjam {
			fmt.Printf("%s - %d \n", p.Item.Name, p.Quantity)
		}
		fmt.Println()
	}

	fmt.Println("\n ==== Overall Loan Summary === \n")

	loanCount := make(map[string]int)

	for _, a := range anggotas {
		for _, p := range a.Pinjam {
			loanCount[p.Item.Name] = loanCount[p.Item.Name] + p.Quantity
		}
	}

	for _, b := range daftarBuku {
		fmt.Printf("%s total borrowed: %d \n", b.Name, loanCount[b.Name])
	}
}
