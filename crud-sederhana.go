package main

import "fmt"

type User struct {
	Email string
	Name  string
	Id    int
}

var users = make(map[int]User)

func CreateUser(id int, name, email string) {
	if _, exist := users[id]; exist {

	}
}

func GetUserById(id int) {
	if user, exist := users[id]; exist {
		fmt.Println(user)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func main() {
	CreateUser(1, "Andi Fandi", "andi@gmail.com")
	CreateUser(2, "Hudan Andi", "huda@gmail.com")
}
