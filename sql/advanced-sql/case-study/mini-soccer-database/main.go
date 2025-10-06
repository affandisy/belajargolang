package main

import (
	"belajar_golang/sql/advanced-sql/case-study/mini-soccer-database/db"
	"database/sql"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Testing")

	database := db.ConnectDB()

	defer database.Close()

	fmt.Println("Please Select your option: ")
	fmt.Println("1. Generate Report")
	fmt.Println("2. List Customer without payment")
	fmt.Println("3. exit")

	var choice int
	fmt.Printf("\n Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		generateRevenueReport(database)
	case 3:
		os.Exit(0)
	}

	// res, err := database.Query("SELECT * FROM customers;")
	// if err != nil {
	// 	fmt.Println("Err: ", err)
	// }

	// fmt.Println("Res: ", res)
}

func generateRevenueReport(db *sql.DB) {
	fmt.Println("Generating Revenue Report...")

	query := `SELECT f.fieldname, COUNT(b.bookingid) as total_booking, SUM(p.paymentamount) AS total_amount FROM bookings b JOIN customers c ON c.customerid = b.customerid JOIN fields f ON f.fieldid = b.fieldid JOIN payments p ON p.bookingid = b.bookingid GROUP BY f.fieldname;`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var fieldname string
		var total_booking float64
		var total_amount float64

		err := rows.Scan(&fieldname, &total_booking, &total_amount)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fieldname, " | ", total_booking, " | ", total_amount)
	}
}
