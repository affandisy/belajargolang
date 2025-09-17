package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

func (b BankAccount) GetBalance() float64 {
	return b.Balance
}

func main() {
	hello := sayHello()

	fmt.Println(hello)

	helloName := HelloWorld("Syihabuddin Affandi")

	fmt.Println(helloName)

	HelloNameAge := HelloNameAge("Syihabuddin Affandi", 25)

	fmt.Println(HelloNameAge)

	total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)

	fmt.Println(total)

	user1 := User{
		Name: "Syihabuddin Affandi",
		Age:  25,
	}

	// Panggil Method
	namaUser := user1.SayHelloMethod()
	fmt.Println(namaUser)

	// Panggil Function
	fmt.Println("Syihabuddin Affandi")

	acc := BankAccount{
		Owner:   "Syihabuddin Affandi",
		Balance: 50000,
	}

	balance := acc.GetBalance()
	fmt.Println("Balance: ", balance)

	acc.Deposit(1000)

	balance = acc.GetBalance()
	fmt.Println("Balance setelahnya:", balance)
}

func sayHello() string {
	return "Hello World"
}

func HelloWorld(name string) string {
	return "Hello, " + name
}

func HelloNameAge(name string, age int) string {
	result := fmt.Sprintf(`Hello, %+v umur kamu %+v`, name, age)
	return result
}

func sum(numbers ...int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return
}

func (u User) SayHelloMethod() (result string) {
	result = fmt.Sprintf(`Hello, %+v from method, Dengan Umur %+v`, u.Name, u.Age)

	return
}
