package infrastructure

import "fmt"

type Mongodb struct {
}

func (m *Mongodb) Select() {
	fmt.Println("Query Select MongoDB")
}
