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
		name    string
		shape   Shape
		hasarea float64
	}{
		{name: "RectAngle", shape: RectAngle{10.0, 20.0}, hasarea: 200.0},
		{name: "Circle", shape: Circle{10.0}, hasarea: 314.0},
		{name: "TriAngle", shape: TriAngle{5.0, 10.0}, hasarea: 25.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.hasarea
			if got != want {
				t.Errorf("%#v got %f want %f", tt.shape, got, want)
			}
		})
		/* 		got := tt.shape.Area()
		   		want := tt.want
		   		if want != got {
		   			t.Errorf("got %f want %f", got, want)
		   		} */
	}
}
