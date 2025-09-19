package adapter

type Database interface {
	Select()
	Insert()
	Update()
	Delete()
}
