package polygons

import (
	"fmt"
	"go_playground/interfaces"
	"reflect"
)

func Calculate(s interfaces.Shape) {
	myType := reflect.TypeOf(s)
	fmt.Println(myType.Name())
	fmt.Println(myType.Kind())

	fmt.Printf("Perimeter: %v\n", s.Perimeter())
	fmt.Printf("Area: %v\n", s.Area())
	if cir, ok := s.(*Circle); ok {
		fmt.Printf("Circle Radius: %v\n", cir.GetRadius())
	}

	// switch x := s.(*type) {
	// 	fmt.Println(x)
	// case Circle:
	// 	fmt.Printf("This is circle %v\n", x)
	// case Rectangle:
	// 	fmt.Println("This is a rectangel")
	// default:
	// 	fmt.Println("This is a polygon")
	// }

	fmt.Println()

}
