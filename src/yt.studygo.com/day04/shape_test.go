package shapes

import "testing"

func TestShape(t *testing.T) {
	rectangle := RectAngle{10.0, 20.0}
	got := Perimeter(rectangle)
	want := 60.0
	if want != got {
		t.Errorf("got %f want %f", got, want)
	}
}
func TestArea(t *testing.T) {
	/* t.Run("rectangles", func(t *testing.T) {
		rectangle := RectAngle{10.0, 20.0}
		got := rectangle.Area()
		want := 200.0
		if want != got {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circles := Circle{10.0}
		got := circles.Area()
		want := 314.0
		if want != got {
			t.Errorf("got %f want %f", got, want)
		}
	}) */

	/* 	// 重构1
	   	CheckArea := func(t *testing.T, shape Shape, want float64) {
	   		t.Helper()
	   		got := shape.Area()
	   		if want != got {
	   			t.Errorf("got %f want %f", got, want)
	   		}
	   	}
	   	t.Run("circles", func(t *testing.T) {
	   		circles := Circle{10.0}
	   		want := 314.0
	   		CheckArea(t, circles, want)
	   	})

	   	t.Run("rectangle", func(t *testing.T) {
	   		rectangle := RectAngle{10.0, 20.0}
	   		want := 200.0
	   		CheckArea(t, rectangle, want)
	   	}) */

	// 重构2 列表驱动测试
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{RectAngle{10.0, 20.0}, 200.0},
		{Circle{10.0}, 314.0},
		{TriAngle{5.0, 10.0}, 25.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		want := tt.want
		if want != got {
			t.Errorf("got %f want %f", got, want)
		}
	}
}
