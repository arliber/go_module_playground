package go_module_playground

import "fmt"

type Circle struct {
	Radius float64
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(%0.2f)", c.Radius)
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
