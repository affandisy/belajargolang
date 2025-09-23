package main

import (
	"fmt"
	"strings"
)

type MenuCoffee struct {
	Name  string
	Price float64
}

type Orders struct {
	Item     MenuCoffee
	Quantity int
}

func main() {
	// Coffe Shop Ordering System

	fmt.Println("Welcome to Hacktiv8 Coffee Shop")

	// Memasukkan data ke struct MenuCoffee
	menu := []MenuCoffee{
		{"Espresso", 3.50},
		{"Cappuccino", 4.00},
		{"Latte", 4.50},
		{"Mocha", 5.00},
		{"Tea", 2.50},
	}

	for _, item := range menu {
		fmt.Printf("%s: $%.2f \n", item.Name, item.Price)
	}

	var order []Orders

	for {
		var penampungItem string
		var penampungQuantity int

		fmt.Print("What would you like to order? (Type 'done' to finish)")
		fmt.Scanln(&penampungItem)

		penampungItem = strings.ToLower(penampungItem)

		if penampungItem == "done" {
			break
		}

		var chosen MenuCoffee
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

		fmt.Printf("How many %s would you like? \n", chosen.Name)
		fmt.Scanln(&penampungQuantity)

		if penampungQuantity <= 0 {
			fmt.Println("Harus lebih dari 0")
			continue
		}

		order = append(order, Orders{Item: chosen, Quantity: penampungQuantity})
	}

	fmt.Println("Your order summary:")

	var subTotal float64

	for _, v := range order {
		totalPriceItem := float64(v.Quantity) * v.Item.Price
		fmt.Printf("%s: %d x $%.2f = $%.2f \n", v.Item.Name, v.Quantity, v.Item.Price, totalPriceItem)
		subTotal = subTotal + totalPriceItem
	}

	fmt.Printf("Subtotal: %.2f \n", subTotal)

	tax := subTotal * 0.10
	tip := subTotal * 0.15

	fmt.Printf("Pajak (10%%): %.2f \n", tax)
	fmt.Printf("Tip (15%%): %.2f \n", tip)

	total := subTotal + tax + tip

	fmt.Printf("Total Akhir: %.2f", total)

}
