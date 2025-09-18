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

	if val, ok := data["Kelas"]; ok {
		fmt.Println("Ada: ", val)
	} else {
		fmt.Println("Tidak Ada")
	}

	dataSlice := make([]map[string]interface{}, 0)

	dataMap1 := map[string]interface{}{
		"Nama":     "Syihabuddin Affandi",
		"Kelas":    8,
		"IsActive": true,
	}

	dataSlice = append(dataSlice, dataMap1)

	dataMap2 := map[string]interface{}{
		"Nama":     "Firmansyah",
		"Kelas":    10,
		"IsActive": false,
	}

	dataSlice = append(dataSlice, dataMap2)

	fmt.Println(dataSlice)

}
