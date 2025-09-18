package main

import "fmt"

// Struct BankAccount
type BankAccountBaru struct {
	Owner   string
	Balance float64
}

// Method Bank DepositBaru
func (b *BankAccountBaru) DepositBaru(amount float64) {
	b.Balance = b.Balance + amount
}

// Method Bank GetBalance
func (b BankAccountBaru) GetBalanceBaru() float64 {
	return b.Balance
}

func main() {
	user1 := BankAccountBaru{
		Owner:   "WS Supratman",
		Balance: 2000,
	}

	fmt.Println("Balance Before: ", user1.Balance)

	user1.DepositBaru(2000)
	fmt.Println("Balance After: ", user1.Balance)

	user1.DepositBaru(3000)
	fmt.Println("Balance After: ", user1.Balance)

}
