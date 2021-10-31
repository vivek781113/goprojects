package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/vivek78113/customdatatypes/organization"
	"github.com/vivek78113/customdatatypes/simplemath"
)

type MathExpr = string

const (
	AddExp   = MathExpr("add")
	Subtract = MathExpr("sub")
	Multiply = MathExpr("mul")
	Divide   = MathExpr("div")
)

func main() {
	// var p organization.Identifiable  = organization.Person{}
	np := organization.NewPerson("Vivek", "Tiwary")
	p := &np
	err := p.SetTwitterhandler("@vivek781113")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.FullName())
	fmt.Println(p.ID())

	var tripper http.RoundTripper = &RoundTripCounter{}
	r, _ := http.NewRequest(http.MethodGet, "http://pluralsight.com", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)
	fmt.Println()

	addExp := EvalExpres(AddExp)
	ar := addExp(2, 3)
	fmt.Println(ar)

	subExp := EvalExpres(Subtract)
	fmt.Println(subExp(89.0, 45.0))

}

type RoundTripCounter struct {
	Count int
}

func (rt *RoundTripCounter) RoundTrip(*http.Request) (*http.Response, error) {
	rt.Count += 1
	return nil, nil
}
func EvalExpres(me MathExpr) func(float64, float64) float64 {
	switch me {
	case AddExp:
		return simplemath.Add
	case Subtract:
		return simplemath.Subtract
	case Multiply:
		return simplemath.Multi
	case Divide:
		return simplemath.Divide
	default:
		return func(f1, f2 float64) float64 {
			return 0
		}
	}
}
