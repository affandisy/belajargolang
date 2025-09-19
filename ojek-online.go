package main

import "fmt"

// Contract Interface Driver
type Driver interface {
	Deliver()
	GetBalance() float64
}

// Struct MotorDriver
type MotorDriver struct {
	Name    string
	Balance float64
}

// Struct CarDriver
type CarDriver struct {
	Name    string
	Balance float64
}

// Struct BoxDriver
type BoxDriver struct {
	Name    string
	Weight  float64
	Balance float64
}

// Method Deliver() MotorDriver
func (m *MotorDriver) Deliver() {
	m.Balance = m.Balance + 10000
	fmt.Println(m.Name, " Telah menyelesaikan pesanan dengan Motor, Saldo anda sekarang, ", m.Balance)
}

// Method GetBalance() MotorDriver
func (m *MotorDriver) GetBalance() float64 {
	return m.Balance
}

// Method Deliver() CarDriver
func (c *CarDriver) Deliver() {
	c.Balance = c.Balance + 20000
	fmt.Println(c.Name, " Telah menyelesaikan pesanan dengan Mobil, Saldo anda sekarang, ", c.Balance)
}

// Method GetBalance() CarDriver
func (c *CarDriver) GetBalance() float64 {
	return c.Balance
}

// Method Deliver() BoxDriver
func (b *BoxDriver) Deliver() {
	b.Balance = b.Balance + 10000*b.Weight
	fmt.Println(b.Name, " Telah menyelesaikan pesanan dengan Box, Saldo anda sekarang, ", b.Balance)
}

// Method GetBalance() BoxDriver
func (b *BoxDriver) GetBalance() float64 {
	return b.Balance
}

type OrderSystem struct{}

func (o OrderSystem) ProcessOrder(d Driver) {
	d.Deliver()
}

func main() {
	motor1 := &MotorDriver{
		Name:    "Syihabuddin Affandi",
		Balance: 0,
	}

	motor2 := &MotorDriver{
		Name:    "Muhammad Firmansyah",
		Balance: 0,
	}

	car1 := &CarDriver{
		Name:    "Muhammad Pascal",
		Balance: 0,
	}

	box1 := &BoxDriver{
		Name:    "Muhammad Rayhan",
		Balance: 0,
		Weight:  2,
	}

	system := OrderSystem{}

	system.ProcessOrder(motor1)
	system.ProcessOrder(motor2)
	system.ProcessOrder(car1)
	system.ProcessOrder(car1)
	system.ProcessOrder(car1)
	system.ProcessOrder(box1)

	// fmt.Println()
}
