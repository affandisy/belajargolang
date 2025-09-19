package infrastructure

import "fmt"

type Postgres struct {
}

func (p *Postgres) Select() {
	fmt.Println("Query Select Postgres")
}
