package model

type AddressData struct {
	Address string
	Note    string
}

type User struct {
	Umur      int
	Nama      string
	IsActive  bool
	Addresses []AddressData
}

func getUser(data User) {
	data.Umur = 10
	data.Nama = "Syihab"
}
