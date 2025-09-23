package main

import (
	"fmt"
	"strings"
)

type CoffeeShop struct {
	Name  string
	Price float64
}

type OrderNew struct {
	Item     CoffeeShop
	Quantity int
}

type CustomerNew struct {
	Name   string
	Orders []OrderNew
}

func main() {
	fmt.Println("Welcome to Coffee Shop")

	// Memasukkan kedalam menu
	menu := []CoffeeShop{
		{"Espresso", 3.50},
		{"Cappuccino", 4.00},
		{"Latte", 5.00},
		{"Mocha", 7.30},
		{"Milkshake", 8.30},
		{"Beer", 9.30},
		{"Chocolate", 3.30},
	}

	for _, v := range menu {
		fmt.Printf("%s: $%.2f \n", v.Name, v.Price)
	}

	var customers []CustomerNew

	for {
		var penampungNama string

		fmt.Print("Masukkan nama anda: (ketik 'done' ketika selesai)")
		fmt.Scanln(&penampungNama)

		// penampungNama = strings.ToLower(penampungNama)

		if penampungNama == "done" {
			break
		}

		customer := CustomerNew{Name: penampungNama}

		fmt.Printf("Welcome %s to Coffe Shop \n", penampungNama)

		for {
			var penampungItem string
			var penampungQuantity int

			fmt.Print("Silahkan order disini: (Type 'done' ketika selesai) \n")
			fmt.Scanln(&penampungItem)

			penampungItem = strings.ToLower(penampungItem)

			if penampungItem == "done" {
				break
			}

			var chosen CoffeeShop
			found := false

			for _, m := range menu {
				if strings.ToLower(m.Name) == penampungItem {
					chosen = m
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Item tidak ada di menu")
				continue
			}

			fmt.Printf("How many %s do you want?", chosen.Name)
			fmt.Scanln(&penampungQuantity)

			if penampungQuantity <= 0 {
				fmt.Println("Harus lebih dari 0")
				continue
			}

			customer.Orders = append(customer.Orders, OrderNew{Item: chosen, Quantity: penampungQuantity})
		}

		customers = append(customers, customer)
	}

	fmt.Println("\n ====== Order Summary per Customer ==== \n")

	for _, c := range customers {
		fmt.Printf("Customer: %s \n", c.Name)
		var subTotal float64
		for _, o := range c.Orders {
			lineTotal := float64(o.Quantity) * o.Item.Price
			fmt.Printf("%s: %d x $%.2f = $%.2f \n", o.Item.Name, o.Quantity, o.Item.Price, lineTotal)
			subTotal = subTotal + lineTotal
		}

		tax := subTotal * 0.10
		tip := subTotal * 0.20
		totalPrice := subTotal + tax + tip

		fmt.Printf("Tax 10%%: %.2f \n", tax)
		fmt.Printf("Tip 20%%: %.2f \n", tip)
		fmt.Printf("Total Price: %.2f \n", totalPrice)
		fmt.Println()
	}

	fmt.Println("\n ============ Sales Summary ========= \n")

	menuCount := make(map[string]int)
	totalRevenue := 0.0

	for _, c := range customers {
		for _, o := range c.Orders {
			menuCount[o.Item.Name] = menuCount[o.Item.Name] + o.Quantity
			totalRevenue = totalRevenue + float64(o.Quantity)*o.Item.Price
		}
	}

	for _, m := range menu {
		fmt.Printf("%s total sold: %d \n", m.Name, menuCount[m.Name])
	}

	fmt.Printf("Total Revenue: $%.2f \n", totalRevenue)

}
