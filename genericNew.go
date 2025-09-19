package main

import "fmt"

type Adder[T NumbersNew] interface {
	Add(a, b T) T
}

type NumbersNew interface {
	int | float32 | float64
}

type IntAdder struct{}

func (IntAdder) Add(a, b int) int {
	return a + b
}

type FloatAdder struct{}

func (FloatAdder) Add(a, b float64) float64 {
	return a + b
}

type MapCRUD[K comparable, V any] struct {
	data map[K]V
}

// Constructor
func NewMapCRUD[K comparable, V any]() *MapCRUD[K, V] {
	return &MapCRUD[K, V]{data: make(map[K]V, 0)}
}

// Upsert Create atau Update (Postgres)
func (m *MapCRUD[K, V]) Save(key K, value V) {
	m.data[key] = value
}

func (m *MapCRUD[K, V]) Get(key K) (V, bool) {
	val, ok := m.data[key]
	return val, ok
}

func (m *MapCRUD[K, V]) Delete(key K) {
	delete(m.data, key)
}

func (m *MapCRUD[K, V]) List() map[K]V {
	return m.data
}

type UserGeneric struct {
	Name  string
	Email string
}

func main() {
	var intAdd Adder[int] = IntAdder{}
	var floatAdd Adder[float64] = FloatAdder{}

	fmt.Println("Int Add: ", intAdd.Add(10, 23))
	fmt.Println("Float Add: ", floatAdd.Add(11.4, 22.57))

	users := NewMapCRUD[string, UserGeneric]()

	users.Save("u1", UserGeneric{Name: "Syihabuddin Affandi", Email: "affandi@mail.com"})

	if val, ok := users.Get("u1"); ok {
		fmt.Println("Name: ", val.Name, " | Email: ", val.Email)
	}

}
