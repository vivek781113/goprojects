package models

type Address struct {
	Road, LandMark, Pin string
}

type Person struct {
	firstName, lastName string
	Add *Address
}

//Getter 
func (p *Person) FirstName() string {
	return p.firstName
}

//Getter
func (p *Person) LastName() string {
	return p.lastName
}

//Setter
func (p *Person) SetName(fn, ln string) {
	p.firstName = fn
	p.lastName = ln
}