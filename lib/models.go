package lib

type Page struct {
	Title       string
	Description string
	Persons     []*Person
}

type Person struct {
	Name    string
	Age     int
	Address string
}
