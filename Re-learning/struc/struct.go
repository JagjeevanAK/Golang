package struc

type Rectangle struct {
	width float64
	height float64
}

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func Perimeter (rectangle Rectangle) float64 {
	return 2 * (rectangle.height + rectangle.width);
}

// Method : same function with different parameters

func (rectangle Rectangle)  Area() float64 {
	return rectangle.height * rectangle.width;
}

func (c Circle) Area() float64{
	return c.radius * 3.14;
}