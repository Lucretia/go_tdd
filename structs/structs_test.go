package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectange := Rectangle{10.0, 10.0}
	got := Perimeter(rectange)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectange := Rectangle{12.0, 6.0}
	got := Area(rectange)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
