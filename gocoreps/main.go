package main

import (
	"fmt"
	"net/http"

	"gocoreps/models"
	"gocoreps/controllers"
)

const (
	c1 = iota
	c3 = 2 << iota
	c2 = iota + 6
	c4
)

type Person struct {
	Id    int
	Name  string
	Marks float32
}

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

	u := models.User{ID: 1, FirstName: "Tom", LastName: "Hardy"}
	fmt.Println(u)

	fmt.Println("**")
	fmt.Println("--*--")
	fmt.Println("**STRUCT DEMO**")
	var u1 Person
	u1.Id = 1
	u1.Name = "Vivek"
	u1.Marks = 8.9
	fmt.Println(u1)
	u2 := Person{Id: 1, Name: "VIVEK", Marks: 61.2}
	fmt.Println(u2)
	fmt.Println("--*--")
	var names [3]string
	names[0] = "Vivek"
	names[1] = "Kumar"
	names[2] = "Tiwary"

	fmt.Println(names)

	nums := [3]int{1, 2, 3}
	fmt.Println(nums)
	fmt.Println(c1, c2, c3, c4)
	fmt.Println("Hello world")

	fmt.Println("**MAP DEMO**")
	m := map[string]int{"foo": 32}
	fmt.Println(m)
	m["foo"] = 23
	fmt.Println(m)
	delete(m, "foo")
	fmt.Println(m)
	fmt.Println("--**--")

	fmt.Println("**SLICE DEMO**")

	slice := nums[:]
	slice[1] = 45
	nums[2] = 47
	fmt.Println(nums, slice)

	sl := []int{0}
	fmt.Println(sl)
	sl = append(sl, 1, 2, 3, 4)
	fmt.Println(sl)

	s1 := sl[1:]
	s2 := sl[:2]
	s3 := sl[1:2]

	fmt.Println(s1, s2, s3)

	fmt.Println("--**--")

	var fn *string = new(string)
	fmt.Println(*fn, fn)
	*fn = "vivek"
	fmt.Println(*fn, fn)

	ln := "Tiwary"
	pt := &ln
	fmt.Println(ln, pt)

	ln = "TIWARY"
	fmt.Println(ln, pt)

	const c int = 3
	fmt.Println(3 + c)
	// fmt.Println(c + 1.2)

}
