package geometry

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := rectangle.CalculatePerimeter()
	expected := 40.0

	if result != expected {
		t.Errorf("expected '%.2f', got '%.2f'", expected, result)
	}
}

func TestArea(t *testing.T) {
	testsArea := []struct {
		form     Shape
		expected float64
	}{
		{Rectangle{10.0, 10.0}, 100.0},
		{Circle{10.0}, math.Pi * math.Pow(10.0, 2)},
		{Triangle{10.0, 5.0}, 25.0},
	}

	for _, tt := range testsArea {
		result := tt.form.CalculateArea()
		if result != tt.expected {
			t.Errorf("expected '%.2f', got '%.2f'", tt.expected, result)
		}
	}
}
