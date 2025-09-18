package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// penggunaan map
	data := map[string]string{}

	data["Nama"] = "Syihabuddin Affandi"
	data["Kelas"] = "6 SD"

	strJSON, _ := json.Marshal(data)
	fmt.Println("map: ", string(strJSON))

	dataInterface := map[string]interface{}{}

	dataInterface["Nama"] = "Dalul"
	dataInterface["IsActive"] = true
	dataInterface["MMR"] = 13000

	strJSON1, _ := json.Marshal(dataInterface)
	fmt.Println("Map Interface: ", string(strJSON1))

	for key, value := range data {
		fmt.Println(key, ": ", value)
	}
}
