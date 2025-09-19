package main

import "fmt"

type UserBaru struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Id    int    `json:"id"`
}

var users = make(map[int]UserBaru)

func CreateUser(id int, name, email string) {
	if _, exist := users[id]; exist {
		fmt.Println("User dengan ID ", id, "Sudah ada")
		return
	}

	users[id] = UserBaru{
		Id:    id,
		Name:  name,
		Email: email,
	}
	fmt.Println("Data Berhasil dibuat")
	return
}

func ListUser() {
	if len(users) == 0 {
		fmt.Println("Belum ada user yang tersimpan")
		return
	}

	for _, user := range users {
		fmt.Println("ID: ", user.Id, "| Email: ", user.Email, "| Name: ", user.Name)
	}

	return
}

func GetUserById(id int) {
	if user, exist := users[id]; exist {
		fmt.Println(user)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func DeleteUser(id int) {
	if _, exist := users[id]; exist {
		delete(users, id)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func UpdateUser(id int, name, email string) {
	if _, exist := users[id]; exist {
		users[id] = UserBaru{Id: id, Name: name, Email: email}
		fmt.Println("Data berhasil di update")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func main() {
	CreateUser(1, "Andi Fandi", "andi@gmail.com")
	CreateUser(2, "Hudan Andi", "huda@gmail.com")
	CreateUser(3, "Nado Demo", "nado@gmail.com")
	CreateUser(4, "Ey Ja", "eyja@gmail.com")
	CreateUser(5, "Xie Shi", "xieshi@gmail.com")

	ListUser()

	GetUserById(1)

	DeleteUser(3)

	ListUser()

	UpdateUser(1, "Data Diupdate", "update@gmail.com")
	ListUser()
}
