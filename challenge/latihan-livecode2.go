package main

import (
	"fmt"
	"strings"
)

type menuCoffes struct {
	Name  string
	Price float64
}

type OrdersNew struct {
	Item     menuCoffes
	Quantity int
}

type Customer struct {
	Name   string
	Orders []OrdersNew
}

func main() {

	menu := []menuCoffes{
		{"Espresso", 3.50},
		{"Cappuccino", 4.00},
		{"Latte", 4.50},
		{"Mocha", 5.00},
		{"Tea", 2.50},
	}

	fmt.Println("Welcome to Hacktiv8 Coffee Shop Upgraded")

	for _, item := range menu {
		fmt.Printf("%s: $%.2f \n", item.Name, item.Price)
	}

	var customers []Customer

	for {
		var penampungNama string

		fmt.Print("Enter Customer Name (or 'done' to finish) : \n")
		fmt.Scanln(&penampungNama)

		if penampungNama == "done" {
			break
		}

		customer := Customer{Name: penampungNama}

		fmt.Printf("Welcome %s, here is the menu... \n", penampungNama)

		for {
			var penampungItem string
			var penampungQuantity int

			fmt.Print("What would you like to order? (or type 'done' to finish order) \n")
			fmt.Scanln(&penampungItem)

			penampungItem = strings.ToLower(penampungItem)

			if penampungItem == "done" {
				break
			}

			var chosen menuCoffes
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

			customer.Orders = append(customer.Orders, OrdersNew{Item: chosen, Quantity: penampungQuantity})
		}

		customers = append(customers, customer)
	}

	fmt.Println("\n====== Order Summary Per Customer ===")

	for _, c := range customers {
		fmt.Printf("Customer: %s \n", c.Name)
		var subTotal float64
		for _, o := range c.Orders {
			lineTotal := float64(o.Quantity) * o.Item.Price
			fmt.Printf("%s: %d x $%.2f = $%.2f \n", o.Item.Name, o.Quantity, o.Item.Price, lineTotal)
			subTotal += lineTotal
		}
		tax := subTotal * 0.10
		tip := subTotal * 0.15
		totalPrice := subTotal + tax + tip

		fmt.Printf("Subtotal: %.2f \n", subTotal)
		fmt.Printf("Tax: %.2f \n", tax)
		fmt.Printf("Tip: %.2f \n", tip)
		fmt.Printf("Total Price: %.2f \n", totalPrice)
	}

	fmt.Println("\n ========= Sales Summary ======= \n")
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

	fmt.Printf("Total revenue: %.2f \n", totalRevenue)

}
