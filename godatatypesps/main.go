package main

import (
	"fmt"

	"github.com/vivek78113/customdatatypes/organization"
)

func main() {
	// var p organization.Identifiable  = organization.Person{}
	p := organization.NewPerson("Vivek", "Tiwary")
	err := p.SetTwitterhandler("@vivek781113")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.FullName())
	fmt.Println(p.ID())
}
