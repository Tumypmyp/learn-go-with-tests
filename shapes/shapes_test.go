package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 2.0, height: 6.0}, hasArea: 12.0},
		{name: "Circle", shape: Circle{radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 12.0, height: 6.0}, hasArea: 3.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f, want %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}

}
