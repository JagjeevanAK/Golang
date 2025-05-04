package struc

type Rectangle struct {
	width int
	height int
}

type Circle struct {
	radius float64
}

func Perimeter (rectangle Rectangle) int {
	return 2 * (rectangle.height + rectangle.width);
}

// Method : same function with different parameters

func (rectangle Rectangle)  Area() int {
	return rectangle.height * rectangle.width;
}

func (c Circle) Area() float64{
	return c.radius * 3.14;
}