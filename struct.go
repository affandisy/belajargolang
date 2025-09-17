package main

import (
	"belajar_golang/model"
	"fmt"
	"unsafe"
)

type AddressData struct {
	Address string
	Note    string
}

type Mahasiswa struct {
	Umur      int
	Nama      string
	isActive  bool
	Addresses []AddressData
}

type Config struct {
	Env              string // 16 byte
	Name             string // 16 byte
	isEnableTimezone bool   // 1 byte
	AuthEnabled      bool   // 1 byte
	Port             int32  // 8 byte
}

func main() {
	mhs1 := model.User{
		Nama:     "Syihab",
		Umur:     10,
		IsActive: true,
		// Addresses: "Jalan Pangeran Darma Kusuma",
	}

	fmt.Println("Nama: ", mhs1.Nama)
	fmt.Println("Umur: ", mhs1.Umur)
	fmt.Println("IsActive?: ", mhs1.IsActive)

	var conf Config
	conf.isEnableTimezone = true
	conf.Env = "prod"
	conf.AuthEnabled = false
	conf.Name = "REST API Syihab"
	conf.Port = 9000

	fmt.Printf("Total Memory Usage Struct: %d \n", unsafe.Sizeof(conf))
	fmt.Println("=========================")
	fmt.Printf("IsEnableTimezone: %d \n", unsafe.Sizeof(conf.isEnableTimezone))
	fmt.Printf("Env: %d \n", unsafe.Sizeof(conf.Env))
	fmt.Printf("AuthEnabled: %d \n", unsafe.Sizeof(conf.AuthEnabled))
	fmt.Printf("Name: %d \n", unsafe.Sizeof(conf.Name))
	fmt.Printf("Port: %d \n", unsafe.Sizeof(conf.Port))

}
