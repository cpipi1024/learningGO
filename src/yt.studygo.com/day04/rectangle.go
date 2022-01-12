package shapes

type Shape interface {
	Area() float64
}

type RectAngle struct {
	height float64
	wide   float64
}

type Circle struct {
	radius float64
}

type TriAngle struct {
	base   float64
	height float64
}

func Perimeter(rectangle RectAngle) float64 {
	return 2 * (rectangle.height + rectangle.wide)
}

// Circle的方法 并且实现了shape接口
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// RectAngle的方法 并且实现了shape接口
func (r RectAngle) Area() float64 {
	return r.height * r.wide
}

// Triangle的方法 并且实现了shape接口
func (t TriAngle) Area() float64 {
	return (t.base * t.height) / 2
}
