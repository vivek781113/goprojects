package polygons

import "math"

type Circle struct {
	Rad float64
}

func (c *Circle) Area() float64 {
	return c.Rad * c.Rad * math.Pi
}
func (c *Circle) Perimeter() float64 {
	return 2 * c.Rad * math.Pi
}

func (c *Circle) GetRadius() float64 {
	return c.Rad
}
