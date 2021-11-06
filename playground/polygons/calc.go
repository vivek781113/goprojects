package polygons

import (
	"fmt"
	"go_playground/interfaces"
)

func Calculate(s interfaces.Shape) {
	fmt.Printf("Area: %v\n", s.Area())
	fmt.Printf("Perimeter: %v\n", s.Perimeter())
}
