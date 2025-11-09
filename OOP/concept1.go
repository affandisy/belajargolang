package main

import "fmt"

// Go menggunakan:
// Structs: Untuk menyimpan data (seperti properties pada class)
// Methods: Fungsi yang "dimiliki" oleh sebuah struct (receiver functions)
// Interfaces: Untuk mendefinisikan perilaku (polimorfisme)
// Composition (Komposisi): Untuk membangun tipe yang kompleks dengan menggabungkan tipe yang lebih kecil

// Keterbatasan level 1
// Hero tidak "memiliki" skill atau item
// Skill tidak bisa "digunakan" untuk mempengaruhi hero lain
// Semua masih terpisah-pisah

// Mendefinisikan "Benda"

// Hero adalah cetak biru dasar untuk hero
type Hero struct {
	Nama string
	HP   int
	Mana int
}

// Attack adalah method yang dimiliki oleh hero
func (h *Hero) Attack(target *Hero) {
	damage := 100
	fmt.Printf("%s menyerang %s, memberikan %d damage! \n", h.Nama, target.Nama, damage)
	target.HP = target.HP - damage
}

// Skill adalah cetak biru untuk hero skill
type Skill struct {
	Name     string
	ManaCost int
	Damage   int
}

// Item adalah cetak biru untuk hero item
type Item struct {
	Nama string
	Cost int
}

func main() {
	lina := &Hero{
		Nama: "Lina",
		HP:   1000,
		Mana: 800,
	}

	axe := &Hero{
		Nama: "Axe",
		HP:   2000,
		Mana: 500,
	}

	lina.Attack(axe)
	fmt.Printf("HP %s sekarang: %d\n", axe.Nama, axe.HP)
}
