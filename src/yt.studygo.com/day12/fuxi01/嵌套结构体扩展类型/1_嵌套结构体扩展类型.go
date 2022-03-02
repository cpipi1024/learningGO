package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y int
}

// 点之间的距离
func (p Point) Distance(q Point) float64 {
	x := (p.X) - (q.X)
	y := (p.Y) - (q.Y)
	distance := math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
	return distance
}

// 放大
func (p *Point) Scale(n int) {
	p.X = n * p.X
	p.Y = n * p.Y
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

type NColoredPoint struct {
	*Point
	Color color.RGBA
}

func main() {
	red := color.RGBA{255, 0, 0, 0}
	blue := color.RGBA{0, 255, 255, 0}

	readpoint := ColoredPoint{Point{1, 1}, red}
	bluepoint := ColoredPoint{Point{5, 4}, blue}

	fmt.Printf("red: %+v \tblue: %+v\n", readpoint, bluepoint)
	distance := readpoint.Distance(bluepoint.Point)
	fmt.Printf("红蓝两点之间的距离:%.2f \n", distance)
	bluepoint.Scale(2)
	fmt.Printf("放大两倍的蓝点:%+v\n", bluepoint)

	// 内嵌的结构体是命名类型指针，这种情况字段和方法会被间接的引入到当前的类型中
	// 添加这一层关系可以共享通用的结构并动态的改变对象之间的关系
	Nredpoint := NColoredPoint{&Point{1, 1}, red}
	Nbluepoint := NColoredPoint{&Point{5, 4}, blue}

	distance1 := Nredpoint.Distance(*Nbluepoint.Point)
	fmt.Printf("两点之间的距离%.2f \n", distance1)

	Nredpoint.Point = Nbluepoint.Point

	fmt.Printf("新红点:%+v\t新蓝点%+v\n", Nredpoint.Point, Nbluepoint.Point)

}
