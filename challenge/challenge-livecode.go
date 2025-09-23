package main

import "fmt"

// Langkah-langkah pengerjaan simulasi Livecode
// Membuat struct
// Mengisi data untuk struct
// Struct dibuat 2 untuk menu restaurant sama order system
// Menggunakan for kosong untuk mengambil data pengguna
// Mendeklarasikan variabel kosong untuk menampung data pengguna

type menus struct {
	Name  string
	Price float64
}

type orders struct {
	Item     menus
	Quantity int
}

func main() {
	menu := []menus{
		{"Pizza", 10.99},
		{"Salad", 7.99},
		{"Burger", 8.99},
		{"Fries", 2.99},
		{"Soda", 1.99},
	}

	fmt.Println("Welcome to Hacktiv8 Restaurant!")
	fmt.Println("Menu: ")

	for _, value := range menu {
		fmt.Printf("%s - $%.2f \n", value.Name, value.Price)
	}

	var order []orders

	for {
		var penampungItem string
		var penampungQuantity int

		fmt.Print("What would you like to order? (Type 'done' to finish) \n")
		fmt.Scanln(&penampungItem)

		if penampungItem == "done" {
			break
		}

		var chosen menus
		found := false

		for _, menu := range menu {
			if menu.Name == penampungItem {
				chosen = menu
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Item tidak ada di menu")
			break
		}

		fmt.Printf("How many %s would you like? \n", penampungItem)
		fmt.Scanln(&penampungQuantity)

		if penampungQuantity <= 0 {
			fmt.Println("Item tersebut harus lebih dari 0")
			continue
		}

		order = append(order, orders{Item: chosen, Quantity: penampungQuantity})
	}

	fmt.Println("\n Your order summary: ")
	var subTotal float64

	for _, o := range order {
		totalPriceItem := float64(o.Quantity) * o.Item.Price
		fmt.Printf("%s: %d x $%.2f = $%.2f \n", o.Item.Name, o.Quantity, o.Item.Price, totalPriceItem)
		subTotal = subTotal + totalPriceItem
	}

	fmt.Printf("Subtotal: $%.2f \n", subTotal)

	taxTotal := subTotal * 0.07
	tipTotal := subTotal * 0.20

	fmt.Printf("Tax (7.00%%): $%.2f \n", taxTotal)
	fmt.Printf("Tip (20.00%%): $%.2f \n", tipTotal)

	totalPrice := subTotal + taxTotal + tipTotal

	fmt.Printf("Total: $%.2f \n", totalPrice)
}
